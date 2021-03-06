// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "strconv"

// Uint16 returns stringer/JSON/text marshaler for the uint16 type.
func Uint16(v uint16) Uint16V { return New().Uint16(v) }

// Uint16 returns stringer/JSON/text marshaler for the uint16 type.
func (Pretty) Uint16(v uint16) Uint16V { return Uint16V{v: v} }

type Uint16V struct {
	v uint16
}

func (v Uint16V) String() string {
	return strconv.FormatUint(uint64(v.v), 10)
}

func (v Uint16V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v Uint16V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
