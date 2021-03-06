// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Int64ps returns stringer/JSON/text marshaler for the int64 pointer slice type.
func Int64ps(a []*int64) Int64PS { return New().Int64ps(a) }

// Int64ps returns stringer/JSON/text marshaler for the int64 pointer slice type.
func (pretty Pretty) Int64ps(a []*int64) Int64PS {
	return Int64PS{
		a:        a,
		prettier: pretty,
	}
}

type Int64PS struct {
	a        []*int64
	prettier Pretty
}

func (a Int64PS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a Int64PS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := Int64p(p).MarshalText()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(a.prettier.separator)
		}
		buf.Write(b)
	}
	return buf.Bytes(), nil
}

func (a Int64PS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Int64p(p).MarshalJSON()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(",")
		}
		buf.Write(b)
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}
