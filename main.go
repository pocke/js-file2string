package main

import (
	"bufio"
	"fmt"
	"io"
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
	r, err := os.Open(fname)
	if err != nil {
		return err
	}

	// XXX: escape
	fmt.Fprintf(w, "exports['%s']='", fname)
	defer w.Write([]byte("';\n"))

	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanBytes)

	for sc.Scan() {
		bs := sc.Bytes()
		switch bs[0] {
		case '\\', '\'':
			w.Write([]byte{'\\'})
			w.Write(bs)
		case '\n':
			w.Write([]byte(`\n`))
		default:
			w.Write(bs)
		}
	}
	return nil
}
