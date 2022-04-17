package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/healeycodes/tools/utils"
)

var (
	DOT       byte = 46
	SEPARATOR      = "  "
	LINEBREAK      = "\n"
)

func main() {
	var hideDotFilesFlag = flag.Bool("h", false, "ignore entries starting with .")
	var detailsFlag = flag.Bool("d", false, "sort by date and display metadata")
	flag.Parse()

	var path string
	if len(flag.Args()) > 0 {
		path = flag.Args()[0]
	} else {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Couldn't get working directory: %s", err)
		}
		path = wd
	}

	display := utils.ListFiles(path, *hideDotFilesFlag, *detailsFlag)
	fmt.Println(display)
}
