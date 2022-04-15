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
	var allFlag = flag.Bool("a", false, "do not ignore entries starting with .")
	var dateFlag = flag.Bool("d", false, "display, and sort by, date")
	flag.Parse()

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	display := utils.ListFiles(path, *allFlag, *dateFlag)
	fmt.Println(display)
}
