package client

import (
	"context"
	"fmt"
	"time"

	"github.com/jonnypillar/somniloquy/config"
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
	Send(*api.UploadRecordRequest) error
	CloseAndRecv() (*api.UploadStatus, error)
}

// Client ...
type Client struct {
	config *config.ClientConfig
	input  Inputter
}

// NewClient ...
func NewClient(c *config.ClientConfig, conn *grpc.ClientConn) (*Client, error) {
	m, err := NewMicrophone(c)
	if err != nil {
		return nil, errors.Wrap(err, "error occurred creating new Client")
	}

	return &Client{
		config: c,
		input:  m,
	}, nil
}

// Stream ...
func (c Client) Stream(ctx context.Context, asc api.RecordServiceClient) error {
	c.input.Start()
	defer c.input.Close()

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			stream, err := asc.Upload(ctx)
			if err != nil {
				errors.Wrap(err, "an error occured creating upload stream")
			}

			err = c.recordMicrophone(ctx, stream)
			if err != nil {
				return errors.Wrap(err, "error occured recording microphone")
			}

			status, err := stream.CloseAndRecv()
			if err != nil {
				return errors.Wrap(err, "failed to receive upstream status response")
			}

			if status.Code != api.UploadStatusCode_Ok {
				return errors.Errorf("failed to upload stream. %s", status.Message)
			}

			fmt.Println("Response from server: ", status)
		}
	}
}

func (c Client) recordMicrophone(ctx context.Context, stream Streamer) error {
	timer := time.NewTimer(c.config.SampleDuration())
	var count int

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Streamer Signalled, len", count)
			return nil
		case <-timer.C:
			fmt.Println("Streamer Signalled, len", count)
			return nil
		default:
			fmt.Println("Sending Chunk")

			req := api.UploadRecordRequest{
				Content: c.input.Read(),
			}

			err := stream.Send(&req)
			if err != nil {
				return errors.Wrap(err, "error occured sending chunk")
			}

			count++
		}
	}
}
