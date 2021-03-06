// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import "bytes"

// Float64s returns stringer/JSON/text marshaler for the float64 slice type.
func Float64s(s []float64) Float64S { return New().Float64s(s) }

// Float64s returns stringer/JSON/text marshaler for the float64 slice type.
func (pretty Pretty) Float64s(s []float64) Float64S {
	return Float64S{
		s:        s,
		prettier: pretty,
	}
}

type Float64S struct {
	s        []float64
	prettier Pretty
}

func (s Float64S) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s Float64S) MarshalText() ([]byte, error) {
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Float64(v).MarshalText()
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

func (s Float64S) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Float64(v).MarshalJSON()
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
