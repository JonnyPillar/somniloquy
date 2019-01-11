package api

import (
	context "context"
	"fmt"
	"io"
	"log"

	"github.com/pkg/errors"
)

// Server ...
type Server struct {
}

// SayHello ...
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Println("Messaged Reseived", in.Greeting)
	return &PingMessage{Greeting: "Bar"}, nil
}

func (s *Server) Upload(stream UploadService_UploadServer) error {
	for {
		c, err := stream.Recv()
		fmt.Println("Stream Received", c)
		if err != nil {
			if err == io.EOF {
				break
			}

			err = errors.Wrapf(err,
				"failed unexpectadely while reading chunks from stream")
			return err
		}
	}

	err := stream.SendAndClose(&UploadStatus{
		Message: "Upload received with success",
		Code:    UploadStatusCode_Ok,
	})
	if err != nil {
		err = errors.Wrapf(err,
			"failed to send status code")
		return err
	}

	log.Println("Messaged Reseived")

	return err
}
