package bdd

import (
	"testing"
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

func TestMultiPrint(t *testing.T) {
	ff := []func(a ...interface{}) string{
		Sputs,
		Sp,
		Sd,
		Sprint,
		Spjson,
	}
	for _, f := range ff {
		if f([]interface{}{1, 2}) != f(1, 2) {
			t.Fail()
		}
	}
}

func TestNewOptional(t *testing.T) {
	o := NewOptional(0, 0, 0)
	o.Print(testdata)
}
