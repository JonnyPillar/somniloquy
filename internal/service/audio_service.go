package service

import (
	"fmt"
	"io"
	"log"

	"github.com/jonnypillar/somniloquy/internal/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// AudioService defines the Audio gRPC Service
type AudioService struct{}

// RegisterAudioService registers the Audio Service with the gRPC Server
func RegisterAudioService(grpcServer *grpc.Server) {
	as := &AudioService{}

	api.RegisterAudioServiceServer(grpcServer, as)
}

// Upload ...
func (s *AudioService) Upload(stream api.AudioService_UploadServer) error {
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

	err := stream.SendAndClose(&api.UploadStatus{
		Message: "Upload received with success",
		Code:    api.UploadStatusCode_Ok,
	})
	if err != nil {
		err = errors.Wrapf(err,
			"failed to send status code")
		return err
	}

	log.Println("Messaged Reseived")

	return err
}
