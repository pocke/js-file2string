package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	if err := Translate(os.Stdin, os.Stdout); err != nil {
		panic(err)
	}
}

func Translate(r io.Reader, w io.Writer) error {
	_, err := w.Write([]byte(`exports.default='`))
	if err != nil {
		return err
	}
	defer w.Write([]byte("'\n"))

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
