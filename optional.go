package bdd

import (
	"fmt"
	"io"
	"reflect"
)

const (
	// Formatted option
	_                  option = 1 << (31 - iota)
	CanDefaultString          // can use .(fmt.Stringer)
	CanFilterDuplicate        // Filter duplicates
	CanRowSpan                // Fold line

	// Formatted style
	_           style = iota
	StylePrint        // Display data; string without quotes
	StyleBPrint       // Display data
	StyleTPrint       // Display type and data
	StyleJPrint       // The json style display; Do not show private
)

type optional struct {
	style style  // Format style
	depth int    // Maximum recursion depth
	opt   option // Option
}

type style int

type option uint32

func NewOptional(depth int, b style, opt option) *optional {
	return &optional{
		style: b,
		opt:   opt,
		depth: depth,
	}
}

func (s *optional) Fprint(w io.Writer, i ...interface{}) (int, error) {
	return fmt.Fprint(w, s.Sprint(i...))
}

func (s *optional) Print(i ...interface{}) (int, error) {
	return fmt.Print(s.Sprint(i...))
}

func (s *optional) Sprint(i ...interface{}) string {
	switch len(i) {
	case 0:
		return ""
	case 1:
		buf := getBuilder()
		defer putBuilder(buf)
		sb := &format{
			buf:      buf,
			filter:   map[uintptr]bool{},
			optional: *s,
		}
		sb.fmt(reflect.ValueOf(i[0]), 0)
		sb.buf.WriteByte('\n')
		ret := sb.buf.String()
		if s.opt.IsCanRowSpan() {
			return Align(ret)
		}
		return ret
	default:
		return s.Sprint(i)
	}
}

func (t option) IsCanDefaultString() bool {
	return (t & CanDefaultString) != 0
}

func (t option) IsCanFilterDuplicate() bool {
	return (t & CanFilterDuplicate) != 0
}

func (t option) IsCanRowSpan() bool {
	return (t & CanRowSpan) != 0
}
