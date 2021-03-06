// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Int8ps returns stringer/JSON/text marshaler for the int8 pointer slice type.
func Int8ps(a []*int8) Int8PS { return New().Int8ps(a) }

// Int8ps returns stringer/JSON/text marshaler for the int8 pointer slice type.
func (pretty Pretty) Int8ps(a []*int8) Int8PS {
	return Int8PS{
		a:        a,
		prettier: pretty,
	}
}

type Int8PS struct {
	a        []*int8
	prettier Pretty
}

func (a Int8PS) String() string {
	b, _ := a.MarshalText()
	return string(b)
}

func (a Int8PS) MarshalText() ([]byte, error) {
	if a.a == nil {
		return []byte(a.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, p := range a.a {
		b, err := a.prettier.Int8p(p).MarshalText()
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

func (a Int8PS) MarshalJSON() ([]byte, error) {
	if a.a == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, p := range a.a {
		b, err := a.prettier.Int8p(p).MarshalJSON()
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
