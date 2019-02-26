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

// DefaultQuietTime ...
const DefaultQuietTime = time.Second

type State int

const (
	Waiting State = iota
	Listening
	Asking
)

type ListenOpts struct {
	State            func(State)
	QuietDuration    time.Duration
	AlreadyListening bool
}

// Inputter ...
type Inputter interface {
	Start()
	Close()
	ReadData() []int32
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
func (c Client) Stream(ctx context.Context, stream Streamer) error {
	err := c.recordMicrophone(ctx, stream)
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
	return nil
}

func (c Client) recordMicrophone(ctx context.Context, stream Streamer) error {
	c.input.Start()
	defer c.input.Close()
	timer := time.NewTimer(c.config.SampleDuration())
	var count int

	results := make([]int32, 64)
	vad := NewVAD(len(results))

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Streamer Signalled, len", count)
			return nil
		case <-timer.C:
			fmt.Println("Streamer Signalled, len", count)
			return nil
		default:
			newData := c.temp(vad)

			fmt.Println("Sending Chunk")

			req := api.UploadRecordRequest{
				Content: newData,
			}

			err := stream.Send(&req)
			if err != nil {
				return errors.Wrap(err, "error occured sending chunk")
			}

			status, err := stream.CloseAndRecv()
			if err != nil {
				return errors.Wrap(err, "failed to receive upstream status response")
			}

			if status.Code != api.UploadStatusCode_Ok {
				return errors.Errorf("failed to upload stream. %s", status.Message)
			}

			fmt.Println("Response from server: ", status)

			count++
		}
	}
}

func (c Client) temp(vad *VAD) []int32 {
	newData := []int32{}
	opts := ListenOpts{}
	heardSomething := opts.AlreadyListening
	quietTime := DefaultQuietTime * 5
	var quietStart time.Time
	var lastFlux float64
	var quiet bool

	//Increasing the level increases the input sensitivy. If it is picking up to much background noise, reduce the value
	level := 0.5

	if opts.State != nil {
		if heardSomething {
			opts.State(Listening)
		} else {
			opts.State(Waiting)
		}
	}
	fmt.Println("Starting")

	for {
		// fmt.Println("Sending Chunk")

		data := c.input.ReadData()

		flux := vad.Flux(data)
		// fmt.Println("Flux Value", flux)
		if lastFlux == 0 {
			lastFlux = flux
			continue
		}

		if heardSomething {
			// fmt.Println("f", flux, "l", lastFlux)

			if flux*level <= lastFlux {
				if !quiet {
					fmt.Println("Resetting queit start")
					quietStart = time.Now()
				} else {
					diff := time.Since(quietStart)

					if diff > quietTime {
						fmt.Println("Breaking Reader", diff, len(newData))
						return newData
					}
				}
				// fmt.Println("quiet")

				quiet = true
			} else {
				newData = append(newData, data...)

				fmt.Println("not quiet")
				quiet = false
				lastFlux = flux
			}
		} else {
			fmt.Println("Not Heard Something")
			if flux >= lastFlux*level {
				heardSomething = true
				if opts.State != nil {
					opts.State(Listening)
				}
			}

			lastFlux = flux
		}
	}
}
