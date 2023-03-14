package bdd

import (
	"fmt"
	"reflect"
)

type format struct {
	optional
	buf    builder
	filter map[uintptr]bool
}

func (s *format) write(b []byte) {
	s.buf.Write(b)
}

func (s *format) writeByte(b byte) {
	s.buf.WriteByte(b)
}

func (s *format) writeString(str string) {
	s.buf.WriteString(str)
}

func (s *format) writeFormat(format string, a ...interface{}) {
	fmt.Fprintf(s.buf, format, a...)
}
