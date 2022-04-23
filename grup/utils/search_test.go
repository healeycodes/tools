package utils

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

var (
	TEST_FILES_DIR = "./finder_test_files/"
)

func TestSearchPathWithLines(t *testing.T) {
	// ----------------------
	rescueStdout := os.Stdout
	r2, w, _ := os.Pipe()
	os.Stdout = w
	// ----------------------

	opts := &SearchOptions{
		LITERAL,
		true,
		nil,
		MakeStringFinder([]byte("c")),
	}
	SearchPath(TEST_FILES_DIR, opts)

	// ------------------------
	w.Close()
	out, _ := ioutil.ReadAll(r2)
	os.Stdout = rescueStdout
	// ------------------------

	want := `finder_test_files/c:1 c
`
	if string(out) != want {
		t.Errorf("TestSearchPath1 was incorrect, got: %s, want: %s", out, want)
	}
}

func TestSearchPathWithoutLines(t *testing.T) {
	// ----------------------
	rescueStdout := os.Stdout
	r2, w, _ := os.Pipe()
	os.Stdout = w
	// ----------------------

	opts := &SearchOptions{
		LITERAL,
		false,
		nil,
		MakeStringFinder([]byte("b")),
	}
	SearchPath(TEST_FILES_DIR, opts)

	// ------------------------
	w.Close()
	out, _ := ioutil.ReadAll(r2)
	os.Stdout = rescueStdout
	// ------------------------

	// Ensure that tests pass on Unix and Windows
	want := `finder_test_files/a/b b
`
	if string(out) != want {
		t.Errorf("TestSearchPath2 was incorrect, got: %s, want: %s", out, want)
	}
}

func TestSearchPathRegexWithLines(t *testing.T) {
	// ----------------------
	rescueStdout := os.Stdout
	r2, w, _ := os.Pipe()
	os.Stdout = w
	// ----------------------

	r, _ := regexp.Compile("b")
	opts := &SearchOptions{
		LITERAL,
		false,
		r,
		nil,
	}
	SearchPath(TEST_FILES_DIR, opts)

	// ------------------------
	w.Close()
	out, _ := ioutil.ReadAll(r2)
	os.Stdout = rescueStdout
	// ------------------------

	want := `finder_test_files/a/b b
`
	if string(out) != want {
		t.Errorf("TestSearchPath2 was incorrect, got: %s, want: %s", out, want)
	}
}
