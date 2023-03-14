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

// getString writes a buffer with the default string,
// returns true if no default string is possible
func (s *format) getString(v reflect.Value) bool {
	if !s.opt.IsCanDefaultString() {
		return false
	}

	switch s.style {
	case StyleJPrint:
		r := getString(v)
		if r == "" {
			return false
		}
		vv, _ := json.Marshal(v.Interface())
		s.write(vv)
	case StyleTPrint:
		r := getString(v)
		if r == "" {
			return false
		}
		s.writeByte('<')
		s.writeString(r)
		s.writeByte('>')
	default:
		r := getString(v)
		if r == "" {
			return false
		}
		s.writeString(r)
	}
	return true
}

// xxBuf writes the buffer in hexadecimal format
func (s *format) xxBuf(v reflect.Value, i interface{}) {
	switch s.style {
	case StyleJPrint:
		s.writeByte('"')
		defer s.writeByte('"')
	case StyleTPrint:
		s.writeByte('<')
		defer s.writeByte('>')
	}
	s.nameBuf(v.Type())
	s.writeFormat("(0x%020x)", i)
	return
}

// getString returns default string
func getString(v reflect.Value) string {
	if v.Kind() == reflect.Interface {
		if v.IsNil() {
			return ""
		}
		return getString(v.Elem())
	}

	if !v.CanInterface() {
		return ""
	}

	i := v.Interface()

	if e, b := i.(fmt.Stringer); b && e != nil {
		return e.String()
	}
	if e, b := i.(fmt.GoStringer); b && e != nil {
		return e.GoString()
	}
	return ""
}
