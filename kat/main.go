package main

import (
	"bufio"
	"log"
	"os"

	"github.com/healeycodes/tools/utils"
)

func main() {
	var r *bufio.Reader
	if len(os.Args) > 1 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		r = bufio.NewReader(f)
	} else {
		info, err := os.Stdin.Stat()
		if err != nil {
			log.Fatal(err)
		}
		if (info.Mode()&os.ModeCharDevice) == 0 || info.Size() <= 0 {
			r = bufio.NewReader(os.Stdin)
		} else {
			log.Fatal("No filepath argument or pipe input")
		}
	}

	utils.StreamFromReader(r)
}
