package client

import (
	"fmt"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/jonnypillar/somniloquy/internal/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Inputter ...
type Inputter interface {
	Start()
	Close()
	Read() []int32
}

// Streamer ...
type Streamer interface {
	Send(*api.UploadAudioRequest) error
	CloseAndRecv() (*api.UploadStatus, error)
}

// Client ...
type Client struct {
	config *config.ClientConfig
	input  Inputter
}

// NewClient ...
func NewClient(config *config.ClientConfig, conn *grpc.ClientConn) (*Client, error) {
	m, err := NewMicrophone(config)
	if err != nil {
		return nil, errors.Wrap(err, "error occurred creating new Client")
	}

	c := Client{
		config: config,
		input:  m,
	}

	return &c, nil
}

// Stream ...
func (c Client) Stream(stream Streamer) error {
	c.input.Start()
	defer c.input.Close()

	shouldSample := true

	// go func() {
	// 	time.Sleep(c.config.SampleDuration())
	// 	fmt.Println("Stopping sampling")

	// 	shouldSample = false
	// }()

	for shouldSample {
		fmt.Println("Sending Chunk")

		req := api.UploadAudioRequest{
			Content: c.input.Read(),
		}

		err := stream.Send(&req)
		if err != nil {
			return errors.Wrap(err, "error occured sending chunk")
		}
		break
	}

	status, err := stream.CloseAndRecv()
	if err != nil {
		return errors.Wrap(err, "failed to receive upstream status response")
	}

	if status.Code != api.UploadStatusCode_Ok {
		return errors.Errorf("failed to upload stream. %s", status.Message)
	}

	fmt.Println("Response from server: ", status)
	return nil
}
