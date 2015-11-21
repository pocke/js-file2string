package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/pflag"
)

type Option struct {
	FileNameOnly bool
}

func main() {
	opt := &Option{}
	pflag.BoolVarP(&opt.FileNameOnly, "filename-only", "f", false, "trim directory")
	pflag.Parse()

	files := pflag.Args()
	if opt.FileNameOnly && !checkFileUniq(files) {
		fmt.Fprintln(os.Stderr, "Files should be uniq")
	}

	for _, fname := range files {
		if err := Translate(fname, os.Stdout, opt); err != nil {
			panic(err)
		}
	}
}

func Translate(fname string, w io.Writer, opt *Option) error {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}

	if opt.FileNameOnly {
		_, fname = filepath.Split(fname)
	}

	// XXX: escape
	fmt.Fprintf(w, "exports['%s']=", fname)
	defer w.Write([]byte(";\n"))

	return json.NewEncoder(w).Encode(string(b))
}

func checkFileUniq(files []string) bool {
	existTable := make(map[string]struct{})
	for _, f := range files {
		_, fname := filepath.Split(f)
		if _, exist := existTable[fname]; exist {
			return false
		}
		existTable[fname] = struct{}{}
	}
	return true
}

func ReplaceFilename(fname string) string {
	var res []byte = nil
	if regexp.MustCompile(`\d`).Match([]byte{fname[0]}) {
		res = make([]byte, 1, len(fname)+1)
		res[0] = '_'
	} else {
		res = make([]byte, 0, len(fname))
	}

	re := regexp.MustCompile(`[[:alnum:]_$]`)
	for _, ch := range []byte(fname) {
		if re.Match([]byte{ch}) {
			res = append(res, ch)
		} else {
			res = append(res, '_')
		}
	}
	return string(res)
}
