package main

import (
	"bufio"
	"log"
	"os"

	"github.com/healeycodes/tools/utils"
)

func main() {
	if len(os.Args) > 1 {
		for _, fp := range os.Args[1:] {
			f, err := os.Open(fp)
			if err != nil {
				log.Fatalf("File error: %s, ", err)
			}
			utils.StreamFromReader(bufio.NewReader(f))
		}
	} else {
		info, err := os.Stdin.Stat()
		if err != nil {
			log.Fatalf("Pipe error: %s", err)
		}
		if (info.Mode()&os.ModeCharDevice) == 0 || info.Size() <= 0 {
			utils.StreamFromReader(bufio.NewReader(os.Stdin))
		} else {
			log.Fatal("No filepath argument or pipe input")
		}
	}
}
