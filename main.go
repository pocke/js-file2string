package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	for _, fname := range os.Args[1:] {
		if err := Translate(fname, os.Stdout); err != nil {
			panic(err)
		}
	}
}

func Translate(fname string, w io.Writer) error {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	// XXX: escape
	fmt.Fprintf(w, "exports['%s']=", fname)
	defer w.Write([]byte(";\n"))

	return json.NewEncoder(w).Encode(string(b))
}
