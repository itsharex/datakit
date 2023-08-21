// @generated Code generated by gen-atomicint.

// Copyright (c) 2020-2023 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package atomic

import (
	"encoding/json"
	"strconv"
	"sync/atomic"
)

// Int32 is an atomic wrapper around int32.
type Int32 struct {
	_ nocmp // disallow non-atomic comparison

	v int32
}

// NewInt32 creates a new Int32.
func NewInt32(val int32) *Int32 {
	return &Int32{v: val}
}

// Load atomically loads the wrapped value.
func (i *Int32) Load() int32 {
	return atomic.LoadInt32(&i.v)
}

// Add atomically adds to the wrapped int32 and returns the new value.
func (i *Int32) Add(delta int32) int32 {
	return atomic.AddInt32(&i.v, delta)
}

// Sub atomically subtracts from the wrapped int32 and returns the new value.
func (i *Int32) Sub(delta int32) int32 {
	return atomic.AddInt32(&i.v, -delta)
}

// Inc atomically increments the wrapped int32 and returns the new value.
func (i *Int32) Inc() int32 {
	return i.Add(1)
}

// Dec atomically decrements the wrapped int32 and returns the new value.
func (i *Int32) Dec() int32 {
	return i.Sub(1)
}

// CAS is an atomic compare-and-swap.
//
// Deprecated: Use CompareAndSwap.
func (i *Int32) CAS(old, new int32) (swapped bool) {
	return i.CompareAndSwap(old, new)
}

// CompareAndSwap is an atomic compare-and-swap.
func (i *Int32) CompareAndSwap(old, new int32) (swapped bool) {
	return atomic.CompareAndSwapInt32(&i.v, old, new)
}

// Store atomically stores the passed value.
func (i *Int32) Store(val int32) {
	atomic.StoreInt32(&i.v, val)
}

// Swap atomically swaps the wrapped int32 and returns the old value.
func (i *Int32) Swap(val int32) (old int32) {
	return atomic.SwapInt32(&i.v, val)
}

// MarshalJSON encodes the wrapped int32 into JSON.
func (i *Int32) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Load())
}

// UnmarshalJSON decodes JSON into the wrapped int32.
func (i *Int32) UnmarshalJSON(b []byte) error {
	var v int32
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	i.Store(v)
	return nil
}

// String encodes the wrapped value as a string.
func (i *Int32) String() string {
	v := i.Load()
	return strconv.FormatInt(int64(v), 10)
}
