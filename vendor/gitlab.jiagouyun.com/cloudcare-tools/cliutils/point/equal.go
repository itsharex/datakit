// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

package point

import (
	"bytes"
	"crypto/md5" //nolint:gosec
	"crypto/sha256"
	"fmt"
	reflect "reflect"
	"sort"
)

// Equal test if two point are the same.
// Equality test not check on warns and debugs.
// If two points equal, they have the same ID(MD5/Sha256),
// but same ID do not means they are equal.
func (p *Point) Equal(x *Point) bool {
	eq, _ := p.EqualWithReason(x)
	return eq
}

func (p *Point) EqualWithReason(x *Point) (bool, string) {
	if x == nil {
		return false, "empty point"
	}

	pname := p.Name() //nolint:ifshort
	ptags := p.Tags()
	pfields := p.Fields()

	xname := x.Name()
	xtags := x.Tags()
	xfields := x.Fields()

	if xtime, ptime := x.Time().UnixNano(), p.Time().UnixNano(); xtime != ptime {
		return false, fmt.Sprintf("timestamp not equal(%d <> %d)", ptime, xtime)
	}

	if !bytes.Equal(xname, pname) {
		return false, fmt.Sprintf("measurement not equla(%s <> %s)", pname, xname)
	}

	if eq, reason := fieldsEqual(pfields, xfields); !eq {
		return eq, reason
	}

	if len(xtags) != len(ptags) {
		return false, fmt.Sprintf("tag count not equal(%d <> %d)", len(ptags), len(xtags))
	}

	for _, t := range ptags {
		if !xtags.KeyExist(t.Key) {
			return false, fmt.Sprintf("tag %s not exists", t.Key)
		}

		v := xtags.Key(t.Key)
		if !bytes.Equal(t.Val, v) {
			return false, fmt.Sprintf("tag %s value not equal(%q <> %q)", t.Key, t.Val, v)
		}
	}

	return true, ""
}

func fieldsEqual(l, r Fields) (bool, string) {
	if len(l) != len(r) {
		return false, fmt.Sprintf("field count not equal(%d <> %d)", len(l), len(r))
	}

	for _, f := range l {
		if !r.KeyExist(f.Key) { // key not exists
			return false, fmt.Sprintf("field %s not exists", string(f.Key))
		}

		v := r.Key(f.Key)
		if !reflect.DeepEqual(f.Val, v) {
			return false, fmt.Sprintf("field %s value not deep equal(%q <> %q)", f.Key, f.Val, v)
		}
	}
	return true, ""
}

// MD5 get point MD5 id.
func (p *Point) MD5() string {
	x := p.hashstr()
	return fmt.Sprintf("%x", md5.Sum(x)) //nolint:gosec
}

// Sha256 get point Sha256 id.
func (p *Point) Sha256() string {
	x := p.hashstr()
	h := sha256.New()
	h.Write(x)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// hashstr only count measurement/tag-keys/tag-values as hash string,
// other fields(fields/time/debugs/warns ignored).
func (p *Point) hashstr() []byte {
	tags := p.Tags()

	var data []byte

	data = append(data, p.Name()...)

	sort.Sort(tags)

	for _, t := range tags {
		data = append(data, t.Key...)
		data = append(data, t.Val...)
	}
	return data
}
