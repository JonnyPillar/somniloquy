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

// Reader ...
type Reader interface {
	//TODO probably not a os.FileInfo from S3
	Read(string) ([]os.FileInfo, error)
}

const languageCode = "en-GB"

// Service ...
type Service struct {
	config *config.ServiceConfig
	reader Reader
}

// NewService registers the Audio Service with the gRPC Server
func NewService(config *config.ServiceConfig, reader Reader) *Service {
	as := Service{
		config: config,
		reader: reader,
	}

	return &as
}

// Start ...
func (ts Service) Start() (Results, error) {
	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a new Google Cloud Speech client")
	}

	flacFiles, err := ts.reader.Read(ts.config.FLACRecordingFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "error occured reading flac recording dir")
	}

	config := speechpb.RecognitionConfig{
		Encoding:        speechpb.RecognitionConfig_FLAC,
		SampleRateHertz: int32(ts.config.SampleRate),
		LanguageCode:    languageCode,
	}

	results := Results{}
	for _, f := range flacFiles {
		if !isFlac(f) {
			continue
		}

		flac := fmt.Sprintf("%s%s", ts.config.FLACRecordingFilePath, f.Name())
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
