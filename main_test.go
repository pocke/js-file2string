package main

import "testing"

func TestCheckFileUniq(t *testing.T) {
	// when uniq
	files := []string{"a", "b", "c"}
	if !checkFileUniq(files) {
		t.Error("Should be true, but got false")
	}

	// when not uniq
	files = []string{"a", "a", "c"}
	if checkFileUniq(files) {
		t.Error("Should be false, but got true")
	}

	files = []string{"b/a", "a", "c"}
	if checkFileUniq(files) {
		t.Error("Should be false, but got true")
	}

	files = []string{"b/a", "d/a", "c"}
	if checkFileUniq(files) {
		t.Error("Should be false, but got true")
	}
}

func TestReplaceFilename(t *testing.T) {
	assert := func(target, expected string) {
		got := ReplaceFilename(target)
		if got != expected {
			t.Errorf("Expected: %q, but got %q", expected, got)
		}
	}

	assert("hoge", "hoge")
	assert(".hoge", "_hoge")
	assert("hoge.html", "hoge_html")
	assert("123", "_123")
	assert("$foo", "$foo")
	assert("foo/bar/nya.go", "foo_bar_nya_go")
}

func TestExportedFilename(t *testing.T) {
	assert := func(target string, opt *Option, expected string) {
		got := ExportedFilename(target, opt)
		if got != expected {
			t.Errorf("Expected: %q, but got %q", expected, got)
		}
	}

	assert("foo", &Option{}, "foo")
	assert("foo/bar", &Option{}, "foo/bar")
	assert("foo/bar", &Option{FileNameOnly: true}, "bar")
	assert("foo/bar", &Option{Replace: true}, "foo_bar")
	assert("foo/bar.html", &Option{FileNameOnly: true, Replace: true}, "bar_html")
}
