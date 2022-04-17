package utils

import (
	"bufio"
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

var (
	TEST_FILES_DIR = "./finder_test_files/"
)

func TestSearch(t *testing.T) {
	reg, _ := regexp.Compile("c")

	// ----------------------
	rescueStdout := os.Stdout
	r2, w, _ := os.Pipe()
	os.Stdout = w
	// ----------------------

	f, _ := os.Open(TEST_FILES_DIR + "c")
	Search(reg, bufio.NewReader(f), &Opts{})

	// ------------------------
	w.Close()
	out, _ := ioutil.ReadAll(r2)
	os.Stdout = rescueStdout
	// ------------------------

	want := `c
`
	if string(out) != want {
		t.Errorf("TestSearch was incorrect, got: %s, want: %s", out, want)
	}
}

func TestRecurse(t *testing.T) {
	reg, _ := regexp.Compile("b")

	// ----------------------
	rescueStdout := os.Stdout
	r2, w, _ := os.Pipe()
	os.Stdout = w
	// ----------------------

	Recurse(reg, TEST_FILES_DIR)

	// ------------------------
	w.Close()
	out, _ := ioutil.ReadAll(r2)
	os.Stdout = rescueStdout
	// ------------------------

	// Ensure that tests pass on Unix and Windows
	wantUnix := `finder_test_files/a/b:1 b
`
	wantWin := `finder_test_files\a\b:1 b
`
	if string(out) != wantUnix && string(out) != wantWin {
		t.Errorf("TestRecurse was incorrect, got: %s, want: %s", out, wantUnix)
	}
}
