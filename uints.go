// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Uints returns stringer/JSON/text marshaler for the uint slice type.
func Uints(s []uint) UintS { return New().Uints(s) }

// Uints returns stringer/JSON/text marshaler for the uint slice type.
func (pretty Pretty) Uints(s []uint) UintS {
	return UintS{
		s:        s,
		prettier: pretty,
	}
}

type UintS struct {
	s        []uint
	prettier Pretty
}

func (s UintS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s UintS) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Uint(v).MarshalText()
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

func (s UintS) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Uint(v).MarshalJSON()
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
