package bdd

import (
	"time"
)

type bbb struct {
	A int
}

type T struct {
	bbb
	Msg  string
	Msg2 string
	Msg3 string
	msg  string
	Msgs []string
	Stru []struct {
		Msg string
		AA  [8]int
	}
	Floats  [6]float32
	Ints    [][]int
	Maps    map[string]string
	B       bool
	T       time.Time
	TTT     interface{}
	TTT2    interface{}
	Chan    interface{}
	Fun     interface{}
	Uintptr uintptr
	Self    *T
}

var testdata *T
