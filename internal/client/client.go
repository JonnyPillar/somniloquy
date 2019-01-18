package client

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/jonnypillar/somniloquy/internal/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Client ...
type Client struct {
	config *config.ClientConfig
	as     api.AudioServiceClient
}

// NewClient ...
func NewClient(config *config.ClientConfig, conn *grpc.ClientConn) *Client {
	asc := api.NewAudioServiceClient(conn)

	c := Client{
		config: config,
		as:     asc,
	}

	return &c
}

// Send ...
func (c Client) Send() error {
	chunks := c.data()

	stream, err := c.as.Upload(context.Background())
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	for i, chunk := range chunks {
		fmt.Println("Sending Chunk", i)
		err = stream.Send(&api.UploadAudioRequest{
			Content: chunk,
		})
	}

	status, err := stream.CloseAndRecv()
	if err != nil {
		return errors.Wrapf(err, "failed to receive upstream status response")
	}

	if status.Code != api.UploadStatusCode_Ok {
		return errors.Errorf("upload failed - msg: %s", status.Message)
	}

	fmt.Println("Response from server: ", status)
	return nil
}

func (c Client) data() [][]byte {
	dat, err := os.Open("./test/data/test.mp3")
	if err != nil {
		fmt.Println("error occurred retreiving test audio", err)
	}

	chunks, err := BufferStream(c.config, dat)
	if err != nil {
		fmt.Println("error occurred retreiving test audio", err)
	}

	return chunks
}
