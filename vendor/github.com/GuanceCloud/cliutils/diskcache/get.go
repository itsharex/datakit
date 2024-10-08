// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package diskcache

import (
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

// Fn is the handler to eat cache from diskcache.
type Fn func([]byte) error

func (c *DiskCache) switchNextFile() error {
	if c.curReadfile != "" {
		if err := c.removeCurrentReadingFile(); err != nil {
			return fmt.Errorf("removeCurrentReadingFile: %w", err)
		}
	}

	// clear .pos
	if !c.noPos {
		if err := c.pos.reset(); err != nil {
			return err
		}
	}

	// reopen next file to read
	return c.doSwitchNextFile()
}

func (c *DiskCache) skipBadFile() error {
	defer func() {
		droppedBatchVec.WithLabelValues(c.path, reasonBadDataFile).Inc()
	}()

	return c.switchNextFile()
}

// Get fetch new data from disk cache, then passing to fn
// if any error occurred during call fn, the reading data is
// dropped, and will not read again.
//
// Get is safe to call concurrently with other operations and will
// block until all other operations finish.
func (c *DiskCache) Get(fn Fn) error {
	var (
		n, nbytes int
		err       error
	)

	c.rlock.Lock()
	defer c.rlock.Unlock()

	start := time.Now()

	defer func() {
		if uint32(nbytes) != EOFHint {
			getBytesVec.WithLabelValues(c.path).Add(float64(nbytes))

			// get on EOF not counted as a real Get
			getVec.WithLabelValues(c.path).Inc()
			getLatencyVec.WithLabelValues(c.path).Observe(float64(time.Since(start) / time.Microsecond))
		}
	}()

	// wakeup sleeping write file, rotate it for succession reading!
	if time.Since(c.wfdLastWrite) > c.wakeup && c.curBatchSize > 0 {
		wakeupVec.WithLabelValues(c.path).Inc()

		if err = func() error {
			c.wlock.Lock()
			defer c.wlock.Unlock()
			return c.rotate()
		}(); err != nil {
			return err
		}
	}

	if c.rfd == nil {
		if err = c.switchNextFile(); err != nil {
			return err
		}
	}

retry:
	if c.rfd == nil {
		return ErrEOF
	}

	hdr := make([]byte, dataHeaderLen)
	if n, err = c.rfd.Read(hdr); err != nil || n != dataHeaderLen {
		//
		// On bad datafile, just ignore and delete the file.
		//
		if err = c.skipBadFile(); err != nil {
			return err
		}

		goto retry // read next new file to save another Get() calling.
	}

	// how many bytes of current data?
	nbytes = int(binary.LittleEndian.Uint32(hdr[0:]))

	if uint32(nbytes) == EOFHint { // EOF
		if err := c.switchNextFile(); err != nil {
			return fmt.Errorf("switchNextFile: %w", err)
		}

		goto retry // read next new file to save another Get() calling.
	}

	databuf := make([]byte, nbytes)

	if n, err = c.rfd.Read(databuf); err != nil {
		return err
	} else if n != nbytes {
		return ErrUnexpectedReadSize
	}

	if fn == nil {
		goto __updatePos
	}

	if err = fn(databuf); err != nil {
		// seek back
		if !c.noFallbackOnError {
			if _, serr := c.rfd.Seek(-int64(dataHeaderLen+nbytes), io.SeekCurrent); serr != nil {
				return serr
			}

			seekBackVec.WithLabelValues(c.path).Inc()
			goto __end // do not update .pos
		}
	}

__updatePos:
	// update seek position
	if !c.noPos {
		c.pos.Seek += int64(dataHeaderLen + nbytes)
		if derr := c.pos.dumpFile(); derr != nil {
			return derr
		}
	}

__end:
	return err
}
