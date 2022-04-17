package utils

import (
	"testing"
)

var (
	TEST_FILES = "./listfiles_test_files/"
)

func BenchmarkListFiles100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ListFiles(TEST_FILES, true, true)
	}
}

func TestListFilesNoArgs(t *testing.T) {
	// Default
	noArgs := ListFiles(TEST_FILES, false, false)
	want := "_/  .dot  b  c"
	if noArgs != want {
		t.Errorf("ListFiles was incorrect, got: %s, want: %s.", noArgs, want)
	}
}

func TestListFilesHideDotFiles(t *testing.T) {
	// Hide dot files
	all := ListFiles(TEST_FILES, true, false)
	want := "_/  b  c"
	if all != want {
		t.Errorf("TestListFilesHideDotFiles was incorrect, got: %s, want: %s.", all, want)
	}
}
