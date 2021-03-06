// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Uint8s returns stringer/JSON/text marshaler for the uint8 slice type.
func Uint8s(s []uint8) Uint8S { return New().Uint8s(s) }

// Uint8s returns stringer/JSON/text marshaler for the uint8 slice type.
func (pretty Pretty) Uint8s(s []uint8) Uint8S {
	return Uint8S{
		s:        s,
		prettier: pretty,
	}
}

type Uint8S struct {
	s        []uint8
	prettier Pretty
}

func (s Uint8S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s Uint8S) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Uint8(v).MarshalText()
		if err != nil {
			return nil, err
		}
		if i != 0 {
			buf.WriteString(s.prettier.separator)
		}
		buf.Write(b)
	}
	return buf.Bytes(), nil
}

func (s Uint8S) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Uint8(v).MarshalJSON()
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
