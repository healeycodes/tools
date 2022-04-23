package main

import (
	"flag"
	"log"
	"regexp"

	"github.com/healeycodes/tools/utils"
)

func parseArgs(flagArgs []string) (string, string) {
	if len(flagArgs) < 2 {
		log.Fatal("usage: fgrup <query> <path>")
	}
	return flagArgs[0], flagArgs[1]
}

func main() {
	var lines = flag.Bool("n", false, "display line number for non-binary files")
	var regex = flag.Bool("re", false, "treat query as a regex")
	flag.Parse()
	query, path := parseArgs(flag.Args())

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

	err := utils.SearchPath(path, opts)
	if err != nil {
		log.Fatalf("couldn't search path %s: %s", path, err)
	}
}
