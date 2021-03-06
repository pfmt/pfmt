// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "strconv"

// Int8 returns stringer/JSON/text marshaler for the int8 type.
func Int8(v int8) Int8V { return New().Int8(v) }

// Int8 returns stringer/JSON/text marshaler for the int8 type.
func (Pretty) Int8(v int8) Int8V {
	return Int8V{v: v}
}

type Int8V struct {
	v int8
}

func (v Int8V) String() string {
	return strconv.Itoa(int(v.v))
}

func (v Int8V) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v Int8V) MarshalJSON() ([]byte, error) {
	return v.MarshalText()
}
