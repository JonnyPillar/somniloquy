package client

import (
	"fmt"
	"io"
	"os"
)

// Record ...
func Record() <-chan []byte {
	output := make(chan []byte)
	dat, err := os.Open("./test/data/test.mp3")
	if err != nil {
		fmt.Println("error occurred retreiving test audio", err)
	}

	go bufferFile(dat, output)

	return output
}

// bufferFile ...
func bufferFile(dat io.ReadCloser, output chan []byte) {
	defer dat.Close()
	defer close(output)

	for {
		buf := make([]byte, (1 << 12))

		_, err := dat.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Finished reading input")
				return
			}

			fmt.Println("error occurred outputting audio byte array", err)
			return
		}

		output <- buf
	}
}
