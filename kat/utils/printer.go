package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

var (
	BUFFER_SIZE = 4 * 1024
)

func StreamFromReader(r *bufio.Reader) {
	nBytes, nChunks := 0, 0
	buf := make([]byte, 0, BUFFER_SIZE)

	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]

		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		nChunks++
		nBytes += int(len(buf))
		fmt.Print(string(buf))

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
}
