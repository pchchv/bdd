package bdd

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

const (
	invalidJSON = "null"
	invalid     = "<nil>"
	private     = "<private>"
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

// nilBuf write buffer with nil
func (s *format) nilBuf() {
	switch s.style {
	case StyleJPrint:
		s.writeString(invalidJSON)
	default:
		s.writeString(invalid)
	}
}

// defaultBuf write buffer with default string
func (s *format) defaultBuf(v reflect.Value) {
	s.nameBuf(v.Type())
	switch s.style {
	case StyleTPrint:
		s.writeByte('(')
		s.writeFormat("%#v", v.Interface())
		s.writeByte(')')
	case StyleJPrint:
		d, _ := json.Marshal(v.Interface())
		s.write(d)
	default:
		s.writeFormat("%+v", v.Interface())
	}
	return
}

// nameBuf write buffer with type name
func (s *format) nameBuf(t reflect.Type) bool {
	switch s.style {
	case StyleTPrint:
		switch t.Kind() {
		case reflect.Map:
			s.writeString("map[")
			s.nameBuf(t.Key())
			s.writeByte(']')
			return s.nameBuf(t.Elem())
		case reflect.Slice:
			s.writeString("[]")
			return s.nameBuf(t.Elem())
		case reflect.Array:
			s.writeByte('[')
			s.writeString(strconv.FormatInt(int64(t.Len()), 10))
			s.writeByte(']')
			return s.nameBuf(t.Elem())
		case reflect.Ptr:
			s.writeByte('*')
			return s.nameBuf(t.Elem())
		default:
			if pkg := t.PkgPath(); pkg != "" {
				s.writeString(pkg)
				s.writeByte('.')
			}

			if t.Name() != "" {
				s.writeString(t.Name())
			} else {
				s.writeString(t.String())
			}
			return true
		}
	}
	return false
}
