// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pfmt

import (
	"bytes"
	"time"
)

// Times returns stringer/JSON/text marshaler for the slice of byte slice type.
func Times(s []time.Time) TimeS { return New().Times(s) }

// Times returns stringer/JSON/text marshaler for the slice of byte slice type.
func (pretty Pretty) Times(s []time.Time) TimeS {
	return TimeS{
		s:        s,
		prettier: pretty,
	}
}

type TimeS struct {
	s        []time.Time
	prettier Pretty
}

func (s TimeS) String() string {
	b, _ := s.MarshalText()
	return string(b)
}

func (s TimeS) MarshalText() ([]byte, error) {
	if s.s == nil {
		return []byte(s.prettier.nil), nil
	}
	var buf bytes.Buffer
	for i, v := range s.s {
		b, err := s.prettier.Time(v).MarshalText()
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

func (s TimeS) MarshalJSON() ([]byte, error) {
	if s.s == nil {
		return []byte("null"), nil
	}
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range s.s {
		b, err := s.prettier.Time(v).MarshalJSON()
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
