package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"

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

type ListenOpts struct {
	State            func(State)
	QuietDuration    time.Duration
	AlreadyListening bool
}

const DefaultQuietTime = time.Second

type State int

const (
	Waiting State = iota
	Listening
	Asking
)

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
	in := make([]int16, 8196)

	c.input.Start()
	defer c.input.Close()

	opts := ListenOpts{
		QuietDuration:    1 * time.Second,
		AlreadyListening: true,
	}

	var (
		buf            bytes.Buffer
		heardSomething = opts.AlreadyListening
		quiet          bool
		quietTime      = opts.QuietDuration
		quietStart     time.Time
		lastFlux       float64
	)

	vad := NewVAD(len(in))

	if quietTime == 0 {
		quietTime = DefaultQuietTime
	}

	if opts.State != nil {
		if heardSomething {
			opts.State(Listening)
		} else {
			opts.State(Waiting)
		}
	}

reader:
	for {
		res := c.input.Read()

		err := binary.Write(&buf, binary.LittleEndian, res)
		if err != nil {
			return err
		}

		flux := vad.Flux(res)

		if lastFlux == 0 {
			lastFlux = flux
			continue
		}

		if heardSomething {
			if flux*1.75 <= lastFlux {
				if !quiet {
					quietStart = time.Now()
				} else {
					diff := time.Since(quietStart)

					if diff > quietTime {
						break reader
					}
				}

				quiet = true
			} else {
				quiet = false
				lastFlux = flux
			}
		} else {
			if flux >= lastFlux*1.75 {
				heardSomething = true
				if opts.State != nil {
					opts.State(Listening)
				}
			}

			lastFlux = flux
		}
	}

	// shouldSample := true

	// go func() {
	// 	time.Sleep(c.config.SampleDuration())
	// 	fmt.Println("Stopping sampling")

	// 	shouldSample = false
	// }()

	// for shouldSample {
	// 	fmt.Println("Sending Chunk")

	req := api.UploadAudioRequest{
		Content: &buf.Bytes(),
	}

	err := stream.Send(&req)
	if err != nil {
		return errors.Wrap(err, "error occured sending chunk")
	}
	// }

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
