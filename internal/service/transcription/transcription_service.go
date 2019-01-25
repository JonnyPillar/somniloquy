package transcription

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"cloud.google.com/go/speech/apiv1"
	"github.com/jonnypillar/somniloquy/configs"
	"github.com/pkg/errors"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

const flacDir = "./assets/recordings/flac/"
const languageCode = "en-GB"

// Service ...
type Service struct {
	config *config.ServiceConfig
}

// NewService registers the Audio Service with the gRPC Server
func NewService(config *config.ServiceConfig) *Service {
	as := Service{
		config: config,
	}

	return &as
}

// Start ...
func (ts Service) Start() (Results, error) {
	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create anew Google Cloud Speech client")
	}

	files, err := ioutil.ReadDir(flacDir)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read flac recording dir")
	}

	config := speechpb.RecognitionConfig{
		Encoding:        speechpb.RecognitionConfig_FLAC,
		SampleRateHertz: int32(ts.config.SampleRate),
		LanguageCode:    languageCode,
	}

	results := Results{}
	for _, f := range files {
		if !isFlac(f) {
			continue
		}

		// Reads the audio file into memory.
		flac := fmt.Sprintf("%s%s", flacDir, f.Name())
		data, err := ioutil.ReadFile(flac)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read flac recording")
		}

		audio := speechpb.RecognitionAudio{
			AudioSource: &speechpb.RecognitionAudio_Content{
				Content: data,
			},
		}

		req := speechpb.RecognizeRequest{
			Config: &config,
			Audio:  &audio,
		}

		// Detects speech in the audio file.
		resp, err := client.Recognize(ctx, &req)
		if err != nil {
			return nil, errors.Wrap(err, "error occurred sending recording to Google Cloud Services API")
		}

		results = append(results, newGCSResults(resp))
	}

	return results, nil
}

func isFlac(f os.FileInfo) bool {
	if !f.Mode().IsRegular() {
		return false
	}

	if filepath.Ext(f.Name()) != ".flac" {
		return false
	}

	return true
}
