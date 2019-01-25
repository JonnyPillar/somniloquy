package record

import (
	"io"
	"log"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/jonnypillar/somniloquy/internal/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// RecordingService defines the Record gRPC Service
type RecordingService struct {
	config *config.ServiceConfig
}

// NewRecordingService registers the Record Service with the gRPC Server
func NewRecordingService(config *config.ServiceConfig, grpcServer *grpc.Server) {
	rs := RecordingService{
		config: config,
	}

	api.RegisterRecordServiceServer(grpcServer, &rs)
}

// Upload ...
func (s *RecordingService) Upload(stream api.RecordService_UploadServer) error {
	r := NewRecording()

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
