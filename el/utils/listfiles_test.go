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

func TestListFiles(t *testing.T) {
	// No arguments
	noArgs := ListFiles(TEST_FILES, false, false)
	want := "_/  b  c"
	if noArgs != want {
		t.Errorf("ListFiles was incorrect, got: %s, want: %s.", noArgs, want)
	}

	// List all files (including dot files)
	all := ListFiles(TEST_FILES, true, false)
	want = "_/  .dot  b  c"
	if all != want {
		t.Errorf("ListFiles was incorrect, got: %s, want: %s.", all, want)
	}

	// List files with date preprended
	date := ListFiles(TEST_FILES, false, true)
	want = `2022-04-15 13:05 _/
2022-04-15 13:02 b
2022-04-15 13:02 c`
	if date != want {
		t.Errorf("ListFiles was incorrect, got: %s, want: %s.", date, want)
	}

	// List all files (including dot files) with date preprended
	allDate := ListFiles(TEST_FILES, true, true)
	want = `2022-04-15 13:49 .dot
2022-04-15 13:05 _/
2022-04-15 13:02 b
2022-04-15 13:02 c`
	if allDate != want {
		t.Errorf("ListFiles was incorrect, got: %s, want: %s.", allDate, want)
	}
}
