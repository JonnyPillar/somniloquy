package service

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
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
	// ctx := context.Background()
	// client, err := speech.NewClient(ctx)
	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }

	// r := Recording{
	// 	File: "test_one",
	// }

	r := NewRecording()

	of, err := os.Create(r.File + ".wav")
	if err != nil {
		panic(err)
	}
	enc := wav.NewEncoder(of, s.config.SampleRate, 24, 1, 1)
	buf := &audio.IntBuffer{Format: &audio.Format{
		NumChannels: 1,
		SampleRate:  s.config.SampleRate,
	}, Data: r.Data}

	for {
		c, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}

			return errors.Wrap(err, "failed unexpectadely while reading chunks from stream")
		}

		r.Append(c.Content)

		z := []int{}

		for _, i := range c.Content {
			z = append(z, int(i))
		}

		buf.Data = z

		if err = enc.Write(buf); err != nil {
			fmt.Println("failed to write to the output file -", err)
			os.Exit(1)
		}
	}

	if err = enc.Close(); err != nil {
		fmt.Println("failed to close the encoder stream")
		os.Exit(1)
	}
	of.Close()

	// err := r.Save()
	// if err != nil {
	// 	return errors.Wrapf(err, "failed to save recording")
	// }

	// if err := exec.Command("ffmpeg", "-i", "./assets/recordings/test_one.aiff", "-c:a", "flac", "./assets/recordings/test_one.flac").Run(); err != nil {
	// 	log.Fatalf("FFMPEG Error: %v", err)
	// }

	// // Reads the audio file into memory.
	// data, err := ioutil.ReadFile("./assets/recordings/" + r.File + ".flac")
	// if err != nil {
	// 	log.Fatalf("Failed to read file: %v", err)
	// }

	// ts := NewTranscriptionService(ctx, s.config, client)
	// res, err := ts.Run(data)
	// if err != nil {
	// 	return errors.Wrap(err, "failed to transcribe audio")
	// }

	// fmt.Println("Transcription Results:", res)

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
