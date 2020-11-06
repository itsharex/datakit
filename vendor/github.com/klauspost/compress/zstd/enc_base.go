package zstd

import (
	"fmt"
	"math/bits"

	"github.com/klauspost/compress/zstd/internal/xxhash"
)

type fastBase struct {
	// cur is the offset at the start of hist
	cur int32
	// maximum offset. Should be at least 2x block size.
	maxMatchOff int32
	hist        []byte
	crc         *xxhash.Digest
	tmp         [8]byte
	blk         *blockEnc
	lastDictID  uint32
}

// CRC returns the underlying CRC writer.
func (e *fastBase) CRC() *xxhash.Digest {
	return e.crc
}

// AppendCRC will append the CRC to the destination slice and return it.
func (e *fastBase) AppendCRC(dst []byte) []byte {
	crc := e.crc.Sum(e.tmp[:0])
	dst = append(dst, crc[7], crc[6], crc[5], crc[4])
	return dst
}

// WindowSize returns the window size of the encoder,
// or a window size small enough to contain the input size, if > 0.
func (e *fastBase) WindowSize(size int) int32 {
	if size > 0 && size < int(e.maxMatchOff) {
		b := int32(1) << uint(bits.Len(uint(size)))
		// Keep minimum window.
		if b < 1024 {
			b = 1024
		}
		return b
	}
	return e.maxMatchOff
}

// Block returns the current block.
func (e *fastBase) Block() *blockEnc {
	return e.blk
}

func (e *fastBase) addBlock(src []byte) int32 {
	if debugAsserts && e.cur > bufferReset {
		panic(fmt.Sprintf("ecur (%d) > buffer reset (%d)", e.cur, bufferReset))
	}
	// check if we have space already
	if len(e.hist)+len(src) > cap(e.hist) {
		if cap(e.hist) == 0 {
			l := e.maxMatchOff * 2
			// Make it at least 1MB.
			if l < 1<<20 {
				l = 1 << 20
			}
			e.hist = make([]byte, 0, l)
		} else {
			if cap(e.hist) < int(e.maxMatchOff*2) {
				panic("unexpected buffer size")
			}
			// Move down
			offset := int32(len(e.hist)) - e.maxMatchOff
			copy(e.hist[0:e.maxMatchOff], e.hist[offset:])
			e.cur += offset
			e.hist = e.hist[:e.maxMatchOff]
		}
	}
	s := int32(len(e.hist))
	e.hist = append(e.hist, src...)
	return s
}

// useBlock will replace the block with the provided one,
// but transfer recent offsets from the previous.
func (e *fastBase) UseBlock(enc *blockEnc) {
	enc.reset(e.blk)
	e.blk = enc
}

func (e *fastBase) matchlenNoHist(s, t int32, src []byte) int32 {
	// Extend the match to be as long as possible.
	return int32(matchLen(src[s:], src[t:]))
}

func (e *fastBase) matchlen(s, t int32, src []byte) int32 {
	if debugAsserts {
		if s < 0 {
			err := fmt.Sprintf("s (%d) < 0", s)
			panic(err)
		}
		if t < 0 {
			err := fmt.Sprintf("s (%d) < 0", s)
			panic(err)
		}
		if s-t > e.maxMatchOff {
			err := fmt.Sprintf("s (%d) - t (%d) > maxMatchOff (%d)", s, t, e.maxMatchOff)
			panic(err)
		}
		if len(src)-int(s) > maxCompressedBlockSize {
			panic(fmt.Sprintf("len(src)-s (%d) > maxCompressedBlockSize (%d)", len(src)-int(s), maxCompressedBlockSize))
		}
	}

	// Extend the match to be as long as possible.
	return int32(matchLen(src[s:], src[t:]))
}

// Reset the encoding table.
func (e *fastBase) resetBase(d *dict, singleBlock bool) {
	if e.blk == nil {
		e.blk = &blockEnc{}
		e.blk.init()
	} else {
		e.blk.reset(nil)
	}
	e.blk.initNewEncode()
	if e.crc == nil {
		e.crc = xxhash.New()
	} else {
		e.crc.Reset()
	}
	if (!singleBlock || d.DictContentSize() > 0) && cap(e.hist) < int(e.maxMatchOff*2)+d.DictContentSize() {
		l := e.maxMatchOff*2 + int32(d.DictContentSize())
		// Make it at least 1MB.
		if l < 1<<20 {
			l = 1 << 20
		}
		e.hist = make([]byte, 0, l)
	}
	// We offset current position so everything will be out of reach.
	// If above reset line, history will be purged.
	if e.cur < bufferReset {
		e.cur += e.maxMatchOff + int32(len(e.hist))
	}
	e.hist = e.hist[:0]
	if d != nil {
		// Set offsets (currently not used)
		for i, off := range d.offsets {
			e.blk.recentOffsets[i] = uint32(off)
			e.blk.prevRecentOffsets[i] = e.blk.recentOffsets[i]
		}
		// Transfer litenc.
		e.blk.dictLitEnc = d.litEnc
		e.hist = append(e.hist, d.content...)
	}
}
