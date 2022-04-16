package utils

import (
	"bufio"
	"io/ioutil"
	"os"
	"testing"
)

var (
	TEST_FILE             = "./printer_test_files/a"
	TEST_STRING_FROM_FILE = "a👍"
)

func TestStreamFromReader(t *testing.T) {
	f, err := os.Open(TEST_FILE)
	if err != nil {
		t.Fatal(err)
	}

	r1 := bufio.NewReader(f)

	// ----------------------
	rescueStdout := os.Stdout
	r2, w, _ := os.Pipe()
	os.Stdout = w
	// ----------------------

	StreamFromReader(r1)

	// ------------------------
	w.Close()
	out, _ := ioutil.ReadAll(r2)
	os.Stdout = rescueStdout
	// ------------------------

	if string(out) != TEST_STRING_FROM_FILE {
		t.Errorf("TestStreamFromFile was incorrect, got: %s, want: %s", out, TEST_STRING_FROM_FILE)
	}
}
