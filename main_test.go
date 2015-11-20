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
