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

func TestListFilesWithDate(t *testing.T) {
	// Prepend, and sort by, date
	date := ListFiles(TEST_FILES, false, true)
	want := `2022-04-15 13:49 .dot
2022-04-15 13:05 _/
2022-04-15 13:02 b
2022-04-15 13:02 c`
	if date != want {
		t.Errorf("TestListFilesWithDate was incorrect, got: %s, want: %s.", date, want)
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

func TestListFilesWithDateAndHideDotFiles(t *testing.T) {
	// Hide dot files
	// (and)
	// Prepend, and sort by, date
	allDate := ListFiles(TEST_FILES, true, true)
	want := `2022-04-15 13:05 _/
2022-04-15 13:02 b
2022-04-15 13:02 c`
	if allDate != want {
		t.Errorf("TestListFilesWithDateAndHideDotFiles was incorrect, got: %s, want: %s.", allDate, want)
	}
}
