package main

import (
	"flag"
	"fmt"
	"os"
)

var w = flag.Bool("w", false, "Write the changes to the file")

func init() {
	flag.Usage = func() {
		w := os.Stdout
		fmt.Fprintf(w, "jsonfmt:\n")
		fmt.Fprintf(w, "Usage:\n")
		fmt.Fprintf(w, "    %s [Options] file1 [filen ...]\n", os.Args[0])
		fmt.Fprintf(w, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}
}
