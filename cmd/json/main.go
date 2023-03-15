package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pchchv/bdd"
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

func format(file string, w bool) error {
	var i interface{}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &i)
	if err != nil {
		return err
	}

	ret := bdd.Spjson(i)
	if !w {
		fmt.Print(ret)
		return nil
	}

	err = ioutil.WriteFile(file, []byte(ret), 0666)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}

	for _, file := range args {
		err := format(file, *w)
		if err != nil {
			fmt.Println(err)
			flag.Usage()
			return
		}
	}
}
