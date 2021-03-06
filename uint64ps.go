// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Uint64ps returns stringer/JSON/text marshaler for the uint64 pointer slice type.
func Uint64ps(a []*uint64) Uint64PS { return New().Uint64ps(a) }

// Uint64ps returns stringer/JSON/text marshaler for the uint64 pointer slice type.
func (pretty Pretty) Uint64ps(a []*uint64) Uint64PS {
	return Uint64PS{
		a:        a,
		prettier: pretty,
	}
}

type Uint64PS struct {
	a        []*uint64
	prettier Pretty
}

func (a Uint64PS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a Uint64PS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Uint64p(p).MarshalText()
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

func (a Uint64PS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Uint64p(p).MarshalJSON()
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
