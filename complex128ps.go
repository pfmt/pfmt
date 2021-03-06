// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Complex128ps returns stringer/JSON/text marshaler for the complex128 pointer slice type.
func Complex128ps(a []*complex128) Complex128PS { return New().Complex128ps(a) }

// Complex128ps returns stringer/JSON/text marshaler for the complex128 pointer slice type.
func (pretty Pretty) Complex128ps(a []*complex128) Complex128PS {
	return Complex128PS{
		a:        a,
		prettier: pretty,
	}
}

type Complex128PS struct {
	a        []*complex128
	prettier Pretty
}

func (a Complex128PS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a Complex128PS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Complex128p(p).MarshalText()
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

func (a Complex128PS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Complex128p(p).MarshalJSON()
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
