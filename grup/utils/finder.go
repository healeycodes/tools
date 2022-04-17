package utils

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type Opts struct {
	Fp *string
}

func Open(fp string) *os.File {
	f, err := os.Open(fp)
	if err != nil {
		log.Fatalf("Couldn't read from %s: %s", fp, err)
	}
	return f
}

func Search(reg *regexp.Regexp, r *bufio.Reader, opts *Opts) {
	scanner := bufio.NewScanner(r)
	c := 0
	for scanner.Scan() {
		c++
		line := scanner.Bytes()
		if found := reg.Find(scanner.Bytes()); found != nil {
			output := string(line)
			if opts.Fp != nil {
				output = fmt.Sprintf("%s:%d %s", *opts.Fp, c, line)
			}
			fmt.Println(output)
		}
	}
}

func Recurse(reg *regexp.Regexp, root string) {
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if root == path {
			return nil
		}

		if !d.IsDir() {
			Search(reg, bufio.NewReader(Open(path)), &Opts{&path})
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Couldn't walk %s: %s", root, err)
	}
}
