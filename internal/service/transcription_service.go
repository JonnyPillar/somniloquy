package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"cloud.google.com/go/speech/apiv1"
	"github.com/jonnypillar/somniloquy/configs"
	"github.com/jonnypillar/somniloquy/internal/service/filesystem"
	"github.com/pkg/errors"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

const (
	languageCode = "en-GB"
	flacExt      = ".flac"
)

// TranscriptionService encapsulates the service that transcribes sleep talking recordings
type TranscriptionService struct {
	config *config.ServiceConfig
	reader filesystem.Reader
}

// NewTranscriptionService initialises a Transcription Service
func NewTranscriptionService(config *config.ServiceConfig, reader Reader) *TranscriptionService {
	as := TranscriptionService{
		config: config,
		reader: reader,
	}

	return &as
}

// Start loads any recordings from the source and sends them the GCS Transcription service.Start
// The results of the request are returned on completion
func (ts TranscriptionService) Start() (TranscriptionResults, error) {
	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create anew Google Cloud Speech client")
	}

	files, err := ts.reader.Read()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read flac recording dir")
	}

	config := speechpb.RecognitionConfig{
		Encoding:        speechpb.RecognitionConfig_FLAC,
		SampleRateHertz: int32(ts.config.SampleRate),
		LanguageCode:    languageCode,
	}

	results := TranscriptionResults{}
	for _, f := range files {
		if !isFlac(f) {
			continue
		}

		// Reads the audio file into memory.
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

		results = append(results, NewTranscriptionResult(resp))
	}

	return results, nil
}

func isFlac(f os.FileInfo) bool {
	if !f.Mode().IsRegular() {
		return false
	}

	if filepath.Ext(f.Name()) != flacExt {
		return false
	}

	return true
}
