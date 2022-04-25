package main

import (
	"flag"
	"log"
	"regexp"

	"github.com/healeycodes/tools/utils"
)

func parseArgs(flagArgs []string) (string, []string) {
	if len(flagArgs) < 2 {
		log.Fatal("usage: fgrup <query> <path>")
	}
	paths := []string{flagArgs[1]}
	return flagArgs[0], paths
}

func main() {
	var lines = flag.Bool("n", false, "display line number for non-binary files")
	var regex = flag.Bool("re", false, "treat query as a regex")
	var workers = flag.Int("w", 128, "[debug] set number of search workers")
	flag.Parse()

	query, paths := parseArgs(flag.Args())
	debug := &utils.SearchDebug{
		Workers: *workers,
	}
	var opts *utils.SearchOptions

	if *regex {
		r, err := regexp.Compile(query)
		if err != nil {
			log.Fatalf("Bad regex: %s", err)
		}
		opts = &utils.SearchOptions{
			Kind:   utils.REGEX,
			Lines:  *lines,
			Regex:  r,
			Finder: nil,
		}
	} else {
		opts = &utils.SearchOptions{
			Kind:   utils.LITERAL,
			Lines:  *lines,
			Regex:  nil,
			Finder: utils.MakeStringFinder([]byte(query)),
		}
	}

	utils.Search(paths, opts, debug)
}
