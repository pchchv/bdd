package bdd

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"unicode"
	"unicode/utf8"
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

// structBuf write buffer with struct
func (s *format) structBuf(v reflect.Value, depth int) {
	s.nameBuf(v.Type())
	s.writeByte('{')
	t := v.Type()
	for i := 0; i != t.NumField(); i++ {
		f := t.Field(i)
		v0 := v.Field(i)
		s.depthBuf(depth + 1)
		s.writeString(f.Name)
		s.writeString(colSym)
		if isPrivateName(f.Name) {
			s.writeString(private)
		} else {
			s.fmt(v0, depth+1)
		}
	}
	s.depthBuf(depth)
	s.writeByte('}')
	return
}

// mapBuf write buffer with map
func (s *format) mapBuf(v reflect.Value, depth int) {
	mk := v.MapKeys()
	valueSlice(mk).Sort()
	s.nameBuf(v.Type())
	s.writeByte('{')
	for i := 0; i != len(mk); i++ {
		k := mk[i]
		switch s.style {
		case StyleJPrint:
			if i != 0 {
				s.depthBuf(depth)
				s.writeByte(',')
			} else {
				s.depthBuf(depth + 1)
			}
		default:
			s.depthBuf(depth + 1)
		}
		s.fmt(k, s.depth-1)
		s.writeString(colSym)
		s.fmt(v.MapIndex(k), depth+1)
	}
	s.depthBuf(depth)
	s.writeByte('}')
	return
}

// depthBuf write buffer with depth
func (s *format) depthBuf(i int) {
	s.writeByte('\n')
	for k := 0; k < i; k++ {
		s.writeByte(Space)
	}
}

// stringBuf write buffer with string
func (s *format) stringBuf(v reflect.Value) {
	switch s.style {
	case StyleTPrint:
		s.defaultBuf(v)
	case StyleBPrint, StyleJPrint:
		s.writeString(strconv.Quote(v.String()))
	default:
		s.writeString(v.String())
	}
	return
}

// funcBuf write buffer with func address
func (s *format) funcBuf(v reflect.Value) {
	switch s.style {
	case StyleJPrint:
		s.writeByte('"')
		defer s.writeByte('"')
	case StyleTPrint:
		s.writeByte('<')
		defer s.writeByte('>')
	}
	s.writeString("func(")
	t := v.Type()
	if t.NumIn() != 0 {
		for i := 0; ; {
			s.writeString(t.In(i).String())
			i++
			if i == t.NumIn() {
				break
			}
			s.writeByte(',')
		}
	}
	s.writeString(")(")
	if t.NumOut() != 0 {
		for i := 0; ; {
			s.writeString(t.Out(i).String())
			i++
			if i == t.NumOut() {
				break
			}
			s.writeByte(',')
		}
	}
	s.writeByte(')')
	s.writeFormat("(0x%020x)", v.Pointer())
	return
}

// sliceBuf write buffer with slice
func (s *format) sliceBuf(v reflect.Value, depth int) {
	s.nameBuf(v.Type())
	s.writeByte('[')
	for i := 0; i != v.Len(); i++ {
		switch s.style {
		case StyleJPrint:
			if i != 0 {
				s.depthBuf(depth)
				s.writeByte(',')
			} else {
				s.depthBuf(depth + 1)
			}
		default:
			s.depthBuf(depth + 1)
		}
		s.fmt(v.Index(i), depth+1)
	}
	s.depthBuf(depth)
	s.writeByte(']')
	return
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

// isPrivateName returns it is a private name
func isPrivateName(name string) bool {
	ch, _ := utf8.DecodeRuneInString(name)
	return !unicode.IsUpper(ch)
}

// struct2Map returns map from struct
func struct2Map(v reflect.Value) map[string]interface{} {
	t := v.Type()
	data := map[string]interface{}{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if isPrivateName(f.Name) {
			continue
		}
		data[f.Name] = v.Field(i).Interface()
	}
	return data
}
