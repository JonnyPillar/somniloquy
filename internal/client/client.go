package client

import (
	"context"
	"fmt"
	"log"

	"github.com/jonnypillar/somniloquy/internal/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Configurer ...
type Configurer interface{}

// Client ...
type Client struct {
	config Configurer
	as     api.AudioServiceClient
}

// NewClient ...
func NewClient(config Configurer, conn *grpc.ClientConn) *Client {
	asc := api.NewAudioServiceClient(conn)

	c := Client{
		config: config,
		as:     asc,
	}

	return &c
}

// Send ...
func (c Client) Send() {
	input := Record()

	stream, err := c.as.Upload(context.Background())
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	for i := range input {
		fmt.Println("Sending Chunk")
		err = stream.Send(&api.UploadAudioRequest{
			Content: i,
		})
	}

	status, err := stream.CloseAndRecv()
	if err != nil {
		err = errors.Wrapf(err,
			"failed to receive upstream status response")
		return
	}

	if status.Code != api.UploadStatusCode_Ok {
		err = errors.Errorf(
			"upload failed - msg: %s",
			status.Message)
		return
	}

	fmt.Println("Response from server: ", status)
}
