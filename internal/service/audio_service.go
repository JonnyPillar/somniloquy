package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"cloud.google.com/go/speech/apiv1"
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
	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	r := Recording{
		File: "test_one",
	}

	// r := NewRecording()

	// for {
	// 	c, err := stream.Recv()

	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}

	// 		return errors.Wrap(err, "failed unexpectadely while reading chunks from stream")
	// 	}

	// 	r.Append(c.Content)
	// }

	// err = r.Save()
	// if err != nil {
	// 	return errors.Wrapf(err, "failed to save recording")
	// }

	// if err := exec.Command("ffmpeg", "-i", "./assets/recordings/test_one.aiff", "-c:a", "flac", "./assets/recordings/test_one.flac").Run(); err != nil {
	// 	log.Fatalf("FFMPEG Error: %v", err)
	// }

	// Reads the audio file into memory.
	data, err := ioutil.ReadFile("./assets/recordings/" + r.File + ".flac")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	ts := NewTranscriptionService(ctx, s.config, client)
	res, err := ts.Run(data)
	if err != nil {
		return errors.Wrap(err, "failed to transcribe audio")
	}

	fmt.Println("Transcription Results:", res)

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
