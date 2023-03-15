package bdd

import "io"

var (
	colSym = ": "
	// Space rune
	Space byte = ' '
	// p go style display types for debug
	defD = NewOptional(10, StyleTPrint, CanFilterDuplicate|CanRowSpan)
	// Pjson json style
	defPjson = NewOptional(20, StyleJPrint, CanDefaultString|CanRowSpan)
	// P go style display types
	defP = NewOptional(5, StyleTPrint, CanDefaultString|CanFilterDuplicate|CanRowSpan)
	// Puts go style
	defPuts = NewOptional(5, StyleBPrint, CanDefaultString|CanFilterDuplicate|CanRowSpan)
	// Print go style
	defPrint = NewOptional(5, StylePrint, CanDefaultString|CanFilterDuplicate|CanRowSpan)
)

// D for debug
func D(a ...interface{}) {
	MarkStack(1, defD.Sprint(a...))
}

// Sd for debug
func Sd(a ...interface{}) string {
	return SmarkStack(1, defD.Sprint(a...))
}

// Fp go style friendly display types and data to writer
func Fp(w io.Writer, a ...interface{}) (int, error) {
	return defP.Fprint(w, a...)
}

// P go style friendly display types and data
func P(a ...interface{}) (int, error) {
	return defP.Print(a...)
}

// Sp go style friendly display types and data to string
func Sp(a ...interface{}) string {
	return defP.Sprint(a...)
}

// Fputs go style friendly to writer
func Fputs(w io.Writer, a ...interface{}) (int, error) {
	return defPuts.Fprint(w, a...)
}

// Puts go style friendly display
func Puts(a ...interface{}) (int, error) {
	return defPuts.Print(a...)
}

// Sputs go style friendly to string
func Sputs(a ...interface{}) string {
	return defPuts.Sprint(a...)
}

// Fprint go style friendly to writer
func Fprint(w io.Writer, a ...interface{}) (int, error) {
	return defPrint.Fprint(w, a...)
}

// Print go style friendly display
func Print(a ...interface{}) (int, error) {
	return defPrint.Print(a...)
}

// Sprint go style friendly to string
func Sprint(a ...interface{}) string {
	return defPrint.Sprint(a...)
}

// Fpjson json style friendly display to writer
func Fpjson(w io.Writer, a ...interface{}) (int, error) {
	return defPjson.Fprint(w, a...)
}

// Pjson json style friendly display
func Pjson(a ...interface{}) (int, error) {
	return defPjson.Print(a...)
}

// Spjson json style friendly display to string
func Spjson(a ...interface{}) string {
	return defPjson.Sprint(a...)
}
