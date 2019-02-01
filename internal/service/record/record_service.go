package record

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/jonnypillar/somniloquy/internal/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// Saver ...
type Saver interface {
	Save(string, *bytes.Buffer) error
}

// RecordingService defines the Record gRPC Service
type RecordingService struct {
	config *config.ServiceConfig
	savers []Saver
}

// NewRecordingService registers the Record Service with the gRPC Server
func NewRecordingService(config *config.ServiceConfig, grpcServer *grpc.Server, savers ...Saver) {
	rs := RecordingService{
		config: config,
		savers: savers,
	}

	api.RegisterRecordServiceServer(grpcServer, &rs)
}

// Upload ...
func (s *RecordingService) Upload(stream api.RecordService_UploadServer) error {
	r := NewAiffEncoder()

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

	b, err := r.Encode()
	if err != nil {
		return errors.Wrapf(err, "failed to create recording buffer")
	}

	fn := aiffFileName()
	for _, saver := range s.savers {
		saver.Save(fn, b)
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

func aiffFileName() string {
	return fmt.Sprintf("%s.%s", time.Now().Format("2006-01-02 15:04:05"), aiffExt)
}
