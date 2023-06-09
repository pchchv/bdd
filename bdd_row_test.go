package bdd

import (
	"bytes"
	"strings"
	"testing"
)

func newRowTestData(l int) [][]int {
	rowtestdata := make([][]int, 0, l)
	for i := 0; i != l; i++ {
		rowtestdata = append(rowtestdata, make([]int, i))
	}
	return rowtestdata
}

func TestRowPrint(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	rowtestdata := newRowTestData(100)
	for _, v := range rowtestdata {
		Fprint(buf, v)
	}

	if strings.TrimSpace(buf.String()) != strings.TrimSpace(rowtestdataout) {
		t.Fail()
	}
}

var rowtestdataout = `
[
]
[
 0
]
[
 0
 0
]
[
 0
 0
 0
]
[
 0 0
 0 0
]
[
 0
 0
 0
 0
 0
]
[
 0 0 0
 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0
 0 0 0 0
]
[
 0 0 0
 0 0 0
 0 0 0
]
[
 0 0 0 0 0
 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0
 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
]
[
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
]
[
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
]
[
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
]
[
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
]
[
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
]
[
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
]
[
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
]
[
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
]
[
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
]
[
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
]
[
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
]
[
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
]
[
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
]
[
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
]
[
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
]
[
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
]
[
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
]
[
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
]
[
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
]
[
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
]
[
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
]
[
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
]
[
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
]
[
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
]
[
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
]
[
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
]
[
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
]
[
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
]
[
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
]
[
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
]
[
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
]
[
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
]
[
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
]
[
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
]
[
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
]
[
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
]
[
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
]
[
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
]
[
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
]
[
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
]
[
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
]
[
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
 0 0 0 0
]
[
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
 0 0 0
]
[
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
 0 0
]
[
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
]
[
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
]
[
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
]
`
