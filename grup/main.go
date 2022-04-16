package main

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/healeycodes/tools/utils"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing query argument")
	}

	r, err := regexp.Compile(os.Args[1])
	if err != nil {
		log.Fatalf("Bad regex: %s", err)
	}

	if len(os.Args) > 2 {
		for _, fp := range os.Args[2:] {
			utils.Recurse(r, fp)
		}
		return
	} else {
		info, err := os.Stdin.Stat()
		if err != nil {
			log.Fatalf("Couldn't stat stdin: %s", err)
		}
		if (info.Mode()&os.ModeCharDevice) == 0 || info.Size() <= 0 {
			utils.Search(r, bufio.NewReader(os.Stdin), &utils.Opts{})
		}
	}
}
