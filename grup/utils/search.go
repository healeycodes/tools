package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sync"
)

const (
	LITERAL = iota
	REGEX
)

type SearchOptions struct {
	Kind   int
	Lines  bool
	Regex  *regexp.Regexp
	Finder *stringFinder
}

type searchJob struct {
	paths []string
	opts  *SearchOptions
	wg    *sync.WaitGroup
}

func searchWorker(searchJobs chan *searchJob, completedJobs chan struct{}) {
	for j := range searchJobs {
		for _, path := range j.paths {
			SearchPath(j.wg, path, j.opts, searchJobs, completedJobs)
			completedJobs <- struct{}{}
		}
	}
}

func Search(paths []string, opts *SearchOptions) {
	searchJobs := make(chan *searchJob)
	completedJobs := make(chan struct{}, 16)

	for w := 0; w < 16; w++ {
		go searchWorker(searchJobs, completedJobs)
	}

	var wg sync.WaitGroup

	completedJobs <- struct{}{}
	wg.Add(1)
	searchJobs <- &searchJob{
		paths,
		opts,
		&wg,
	}
	wg.Wait()
}

func SearchPath(wg *sync.WaitGroup, path string, opts *SearchOptions, searchJobs chan *searchJob, completedJobs chan struct{}) error {
	defer wg.Done()
	info, err := os.Lstat(path)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		return SearchFile(path, opts)
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}
	dirNames, err := f.Readdirnames(-1)
	if err != nil {
		return err
	}

	paths := make([]string, len(dirNames))
	for i, dirPath := range dirNames {
		paths[i] = filepath.Join(path, dirPath)
	}

	wg.Add(len(dirNames))
	<-completedJobs
	searchJobs <- &searchJob{
		paths,
		opts,
		wg,
	}
	return nil
}

func SearchFile(path string, opts *SearchOptions) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(f)
	isBinary := false

	line := 1
	for scanner.Scan() {
		text := scanner.Bytes()
		if line == 1 {
			isBinary = bytes.IndexByte(text, 0) != -1
		}

		if opts.Kind == LITERAL {
			if opts.Finder.next(text) != -1 {
				if isBinary {
					fmt.Printf("Binary file %s matches\n", path)
					return err
				} else if opts.Lines {
					fmt.Printf("%s:%d %s\n", path, line, text)
				} else {
					fmt.Printf("%s %s\n", path, text)
				}
			}
		} else if opts.Kind == REGEX {
			if opts.Regex.Find(scanner.Bytes()) != nil {
				if isBinary {
					fmt.Printf("Binary file %s matches\n", path)
					return err
				} else if opts.Lines {
					fmt.Printf("%s:%d %s\n", path, line, text)
				} else {
					fmt.Printf("%s %s\n", path, text)
				}
			}
		}
		line++
	}
	return nil
}

// Below, is Go's internal Boyer-Moore string search algorithm, it has been
// modified to use []byte instead of string to reduce allocations.

// https://go.googlesource.com/go/+/go1.18.1/src/strings/search.go
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// stringFinder efficiently finds strings in a source text. It's implemented
// using the Boyer-Moore string search algorithm:
// https://en.wikipedia.org/wiki/Boyer-Moore_string_search_algorithm
// https://www.cs.utexas.edu/~moore/publications/fstrpos.pdf (note: this aged
// document uses 1-based indexing)
type stringFinder struct {
	// pattern is the string that we are searching for in the text.
	pattern []byte

	// badCharSkip[b] contains the distance between the last byte of pattern
	// and the rightmost occurrence of b in pattern. If b is not in pattern,
	// badCharSkip[b] is len(pattern).
	//
	// Whenever a mismatch is found with byte b in the text, we can safely
	// shift the matching frame at least badCharSkip[b] until the next time
	// the matching char could be in alignment.
	badCharSkip [256]int

	// goodSuffixSkip[i] defines how far we can shift the matching frame given
	// that the suffix pattern[i+1:] matches, but the byte pattern[i] does
	// not. There are two cases to consider:
	//
	// 1. The matched suffix occurs elsewhere in pattern (with a different
	// byte preceding it that we might possibly match). In this case, we can
	// shift the matching frame to align with the next suffix chunk. For
	// example, the pattern "mississi" has the suffix "issi" next occurring
	// (in right-to-left order) at index 1, so goodSuffixSkip[3] ==
	// shift+len(suffix) == 3+4 == 7.
	//
	// 2. If the matched suffix does not occur elsewhere in pattern, then the
	// matching frame may share part of its prefix with the end of the
	// matching suffix. In this case, goodSuffixSkip[i] will contain how far
	// to shift the frame to align this portion of the prefix to the
	// suffix. For example, in the pattern "abcxxxabc", when the first
	// mismatch from the back is found to be in position 3, the matching
	// suffix "xxabc" is not found elsewhere in the pattern. However, its
	// rightmost "abc" (at position 6) is a prefix of the whole pattern, so
	// goodSuffixSkip[3] == shift+len(suffix) == 6+5 == 11.
	goodSuffixSkip []int
}

func MakeStringFinder(pattern []byte) *stringFinder {
	f := &stringFinder{
		pattern:        pattern,
		goodSuffixSkip: make([]int, len(pattern)),
	}
	// last is the index of the last character in the pattern.
	last := len(pattern) - 1

	// Build bad character table.
	// Bytes not in the pattern can skip one pattern's length.
	for i := range f.badCharSkip {
		f.badCharSkip[i] = len(pattern)
	}
	// The loop condition is < instead of <= so that the last byte does not
	// have a zero distance to itself. Finding this byte out of place implies
	// that it is not in the last position.
	for i := 0; i < last; i++ {
		f.badCharSkip[pattern[i]] = last - i
	}

	// Build good suffix table.
	// First pass: set each value to the next index which starts a prefix of
	// pattern.
	lastPrefix := last
	for i := last; i >= 0; i-- {
		if bytes.HasPrefix(pattern, pattern[i+1:]) {
			lastPrefix = i + 1
		}
		// lastPrefix is the shift, and (last-i) is len(suffix).
		f.goodSuffixSkip[i] = lastPrefix + last - i
	}
	// Second pass: find repeats of pattern's suffix starting from the front.
	for i := 0; i < last; i++ {
		lenSuffix := longestCommonSuffix(pattern, pattern[1:i+1])
		if pattern[i-lenSuffix] != pattern[last-lenSuffix] {
			// (last-i) is the shift, and lenSuffix is len(suffix).
			f.goodSuffixSkip[last-lenSuffix] = lenSuffix + last - i
		}
	}

	return f
}

func longestCommonSuffix(a, b []byte) (i int) {
	for ; i < len(a) && i < len(b); i++ {
		if a[len(a)-1-i] != b[len(b)-1-i] {
			break
		}
	}
	return
}

// next returns the index in text of the first occurrence of the pattern. If
// the pattern is not found, it returns -1.
func (f *stringFinder) next(text []byte) int {
	i := len(f.pattern) - 1
	for i < len(text) {
		// Compare backwards from the end until the first unmatching character.
		j := len(f.pattern) - 1
		for j >= 0 && text[i] == f.pattern[j] {
			i--
			j--
		}
		if j < 0 {
			return i + 1 // match
		}
		i += max(f.badCharSkip[text[i]], f.goodSuffixSkip[j])
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
