# **bdd** [![Go Report Card](https://goreportcard.com/badge/github.com/pchchv/bdd)](https://goreportcard.com/report/github.com/pchchv/bdd) [![Go Reference](https://pkg.go.dev/badge/github.com/pchchv/bdd.svg)](https://pkg.go.dev/github.com/pchchv/bdd) [![GitHub license](https://img.shields.io/github/license/pchchv/bdd.svg)](https://github.com/pchchv/bdd/blob/master/LICENSE)

### **B**eautiful **d**ata **d**isplay for Go

## Usage

## *Import*
```go
import "github.com/pchchv/bdd"
```

**[Examples](https://github.com/pchchv/bdd/blob/master/examples/main.go)**

```go
package main

import (
	"fmt"

	"github.com/pchchv/bdd"
)

type mt struct {
	String string
	Int    int
	Slice  []int
	Map    map[string]interface{}
}

func main() {
	m := mt{
		"hello world",
		100,
		[]int{1, 2, 3, 4, 5, 6},
		map[string]interface{}{
			"A":  123,
			"BB": 456,
		},
	}

	fmt.Println(m) // fmt Default Output
	/*
		{hello world 100 [1 2 3 4 5 6] map[BB:456 A:123]}
	*/

	bdd.Puts(m) // More friendly output
	/*
		{
		 String: "hello world"
		 Int:    100
		 Slice:  [
		  1 2 3
		  4 5 6
		 ]
		 Map: {
		  "A":  123
		  "BB": 456
		 }
		}
	*/

	bdd.Print(m) // Same as Puts, but without the quotation marks
	/*
		{
		 String: hello world
		 Int:    100
		 Slice:  [
		  1 2 3
		  4 5 6
		 ]
		 Map: {
		  A:  123
		  BB: 456
		 }
		}
	*/

	bdd.P(m) // Friendly formatting plus type
	/*
		main.mt{
		 String: string("hello world")
		 Int:    int(100)
		 Slice:  []int[
		  int(1) int(2) int(3)
		  int(4) int(5) int(6)
		 ]
		 Map: map[string]interface {}{
		  string("A"):  int(123)
		  string("BB"): int(456)
		 }
		}
	*/

	bdd.Pjson(m) // Output in json style
	/*
		{
		 "Int": 100
		,"Map": {
		  "A":  123
		 ,"BB": 456
		 }
		,"Slice": [
		  1,2,3
		 ,4,5,6
		 ]
		,"String": "hello world"
		}
	*/

	m0 := bdd.ToTable(m, m) // Split into tables by field
	bdd.Puts(m0)
	/*
		[
		 [
		  "String" "Int"
		  "Slice"  "Map"
		 ]
		 [
		  "hello world"   "100"
		  "[1 2 3 4 5 6]" "map[A:123 BB:456]"
		 ]
		]
	*/

	m1 := bdd.FmtTable(m0) // [][]string Table Formatting
	bdd.Puts(m1)
	/*
		[
		 "String      Int Slice         Map               "
		 "hello world 100 [1 2 3 4 5 6] map[A:123 BB:456] "
		]
	*/

	bdd.Mark("hello") // Mark output position
	/*
		main.go:124  hello
	*/

	bdd.Print(bdd.BytesViewer("Hello world! Hello All!"))
	/*
		| Address  | Hex                                             | Text             |
		| -------: | :---------------------------------------------- | :--------------- |
		| 00000000 | 48 65 6c 6c 6f 20 77 6f 72 6c 64 21 20 48 65 6c | Hello world! Hel |
		| 00000010 | 6c 6f 20 41 6c 6c 21                            | lo All!          |
	*/
}
```