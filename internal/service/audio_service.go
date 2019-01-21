package service

import (
	"io"
	"log"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/jonnypillar/somniloquy/internal/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// AudioService defines the Audio gRPC Service
type AudioService struct {
	config *config.ServiceConfig
}

// NewAudioService registers the Audio Service with the gRPC Server
func NewAudioService(config *config.ServiceConfig, grpcServer *grpc.Server) {
	as := &AudioService{
		config: config,
	}

	api.RegisterAudioServiceServer(grpcServer, as)
}

// Upload ...
func (s *AudioService) Upload(stream api.AudioService_UploadServer) error {
	r := Recording{}

	for {
		c, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}

			return errors.Wrap(err, "failed unexpectadely while reading chunks from stream")
		}

		r.Append(c.Content)
	}

	err := r.Save()
	if err != nil {
		return errors.Wrapf(err, "failed to save recording")
	}

	status := api.UploadStatus{
		Message: "Upload received with success",
		Code:    api.UploadStatusCode_Ok,
	}

	err = stream.SendAndClose(&status)
	if err != nil {
		return errors.Wrap(err, "failed to send status code")
	}

	log.Println("Upload Successful")

	return nil
}
