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
	Replace      bool
}

func main() {
	opt := &Option{}
	pflag.BoolVarP(&opt.FileNameOnly, "filename-only", "f", false, "trim directory")
	pflag.BoolVarP(&opt.Replace, "replace", "r", false, "replace as javascript identifier")
	pflag.Parse()

	files := pflag.Args()
	if (opt.FileNameOnly || opt.Replace) && !checkFileUniq(files, opt) {
		fmt.Fprintln(os.Stderr, "Files should be uniq")
		os.Exit(1)
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

	efname := ExportedFilename(fname, opt)

	fmt.Fprintf(w, "exports['%s']=", efname)
	defer w.Write([]byte(";\n"))

	bs, err := json.Marshal(string(b))
	if err != nil {
		return err
	}
	w.Write(bs)
	return nil
}

func checkFileUniq(files []string, opt *Option) bool {
	existTable := make(map[string]struct{})
	for _, f := range files {
		fname := ExportedFilename(f, opt)
		if _, exist := existTable[fname]; exist {
			return false
		}
		existTable[fname] = struct{}{}
	}
	return true
}

func ReplaceFilename(fname string) string {
	var res []byte
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

func ExportedFilename(fpath string, opt *Option) string {
	res := fpath
	if opt.FileNameOnly {
		_, res = filepath.Split(fpath)
	}
	if opt.Replace {
		res = ReplaceFilename(res)
	}
	return res
}
