package client

import (
	"fmt"
	"io"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/pkg/errors"
)

// BufferStream ...
func BufferStream(config *config.ClientConfig, source io.Reader) ([][]byte, error) {
	output := [][]byte{}

	for {
		buf := make([]byte, config.StreamChunkSize)

		_, err := source.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("finished reading input")
				break
			}

			return nil, errors.Wrap(err, "error occurred reading byte array")
		}

		output = append(output, buf)
	}

	return output, nil
}
