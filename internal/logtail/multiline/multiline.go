// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package multiline wraps text multiline match functions
package multiline

import (
	"bytes"
	"time"
	"unicode"
	"unicode/utf8"
)

const (
	defaultMaxLength       = 4 * 1024 * 1024
	defaultMaxLifeDuration = time.Second * 5
)

type option struct {
	// 限制一段多行数据的最大长度，避免出现超级长的多行数据，超出限制会执行 flush
	maxLength int

	// 限制一段多行数据的最大存在时长，即从第一条匹配成功开始到现在，超出限制会执行 flush
	// 避免出现一条匹配成功，N 条匹配失败然后追加写入到 buff，导致数据全部堆积的情况
	maxLifeDuration time.Duration
}

type Option func(*option)

func WithMaxLength(max int) Option {
	return func(opt *option) {
		if max > 0 {
			opt.maxLength = max
		}
	}
}

func WithMaxLifeDuration(dur time.Duration) Option {
	return func(opt *option) {
		if dur > 0 {
			opt.maxLifeDuration = dur
		}
	}
}

func defaultOption() *option {
	return &option{
		maxLength:       defaultMaxLength,
		maxLifeDuration: defaultMaxLifeDuration,
	}
}

type Multiline struct {
	*Matcher
	buff bytes.Buffer
	opt  *option

	// 记录最后一次匹配成功并写入到 buff 的时间
	lastWriteTime time.Time
}

func New(patterns []string, opts ...Option) (*Multiline, error) {
	c := defaultOption()
	for _, opt := range opts {
		opt(c)
	}

	match, err := NewMatcher(patterns)
	if err != nil {
		return nil, err
	}

	return &Multiline{
		Matcher: match,
		opt:     c,
	}, err
}

var newLine = []byte{'\n'}

func (m *Multiline) ProcessLineString(text string) (string, State) {
	if m.MatchString(text) {
		finishedText := m.FlushString()
		m.buff.WriteString(text)
		m.lastWriteTime = time.Now()
		return finishedText, NewMultiline
	}

	if m.buff.Len() == 0 {
		return text, NoContext
	}

	m.buff.Write(newLine)
	m.buff.WriteString(text)

	if time.Since(m.lastWriteTime) > m.opt.maxLifeDuration {
		return m.FlushString(), OverTime
	}
	if m.buff.Len() > m.opt.maxLength {
		return m.FlushString(), OverLength
	}

	return "", Written
}

func (m *Multiline) ProcessLine(text []byte) ([]byte, State) {
	// --匹配成功--
	// 清空 buff 并写入新的文本，符合多行行为。记录当前时间。
	if m.Match(text) {
		finishedText := m.Flush()
		m.buff.Write(text)
		m.lastWriteTime = time.Now()
		return finishedText, NewMultiline
	}

	// --匹配失败--

	// 这一条文本匹配失败，原应该追加写入到 buff 中，但是此时 buff 为空，说明这条文本没有头，是一条“僵尸多行文本”
	// 为了避免匹配失败的文本堆积在 buff 中，需要在此处直接 return
	if m.buff.Len() == 0 {
		return text, NoContext
	}

	// buff 不为空，说明 buff 中存在匹配成功的文本
	// 将本次匹配失败的文本写入到 buff，追加到末尾，符合多行行为
	m.buff.Write(newLine)
	m.buff.Write(text)

	// flush 规则一：单次多行采集时长超出限制
	if time.Since(m.lastWriteTime) > m.opt.maxLifeDuration {
		return m.Flush(), OverTime
	}
	// flush 规则二：buff 长度超过限制
	if m.buff.Len() > m.opt.maxLength {
		return m.Flush(), OverLength
	}

	// 不符合 flush 条件，文本数据都在 buff 中
	return nil, Written
}

func (m *Multiline) BuffLength() int {
	return m.buff.Len()
}

func (m *Multiline) Flush() []byte {
	if m.buff.Len() == 0 {
		return nil
	}

	text := make([]byte, m.buff.Len())
	copy(text, m.buff.Bytes())

	m.buff.Reset()
	return text
}

func (m *Multiline) FlushString() string {
	if m.buff.Len() == 0 {
		return ""
	}
	text := m.buff.String()
	m.buff.Reset()
	return text
}

var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

func TrimRightSpace(s []byte) []byte {
	end := len(s)
	for ; end > 0; end-- {
		c := s[end-1]
		if c >= utf8.RuneSelf {
			return bytes.TrimFunc(s[:end], unicode.IsSpace)
		}
		if asciiSpace[c] == 0 {
			break
		}
	}
	return s[:end]
}
