// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Float32ps returns stringer/JSON/text marshaler for the float32 pointer slice type.
func Float32ps(a []*float32) Float32PS { return New().Float32ps(a) }

// Float32ps returns stringer/JSON/text marshaler for the float32 pointer slice type.
func (pretty Pretty) Float32ps(a []*float32) Float32PS {
	return Float32PS{
		a:        a,
		prettier: pretty,
	}
}

type Float32PS struct {
	a        []*float32
	prettier Pretty
}

func (a Float32PS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a Float32PS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Float32p(p).MarshalText()
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

func (a Float32PS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Float32p(p).MarshalJSON()
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
