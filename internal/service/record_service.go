package service

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jonnypillar/somniloquy/config"
	"github.com/jonnypillar/somniloquy/internal/api"
	"github.com/pkg/errors"
)

// Saver ...
type Saver interface {
	Save(string, *bytes.Buffer) error
}

// RecordingService defines the Record gRPC Service
type RecordingService struct {
	config *config.ServiceConfig
	saver  Saver
}

// NewRecordingService registers the Record Service with the gRPC Server
func NewRecordingService(c *config.ServiceConfig, s Saver) *RecordingService {
	rs := RecordingService{
		config: c,
		saver:  s,
	}

	return &rs
}

// Upload ...
func (s *RecordingService) Upload(stream api.RecordService_UploadServer) error {
	r := NewAiffEncoder()
	var count int

	for {
		c, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}

			return errors.Wrap(err, "failed unexpectadely while reading chunks from stream")
		}

		r.Append(c.Content)

		count++
	}

	b, err := r.Encode()
	if err != nil {
		return errors.Wrap(err, "failed to create recording buffer")
	}

	fn := aiffFileName()
	err = s.saver.Save(fn, b)
	if err != nil {
		return errors.Wrap(err, "error occured saving recording")
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
	return fmt.Sprintf("%s%s", time.Now().Format("2006-01-02 15:04:05"), aiffExt)
}
