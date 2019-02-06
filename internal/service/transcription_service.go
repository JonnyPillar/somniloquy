package service

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jonnypillar/somniloquy/config"
	"github.com/jonnypillar/somniloquy/internal/service/filesystem"
	"github.com/pkg/errors"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

const (
	languageCode = "en-GB"
	flacExt      = ".flac"
)

// SpeechRecogniser ...
type SpeechRecogniser interface {
	Recognize(context.Context, *speechpb.RecognizeRequest) (*TranscriptionResult, error)
}

// TranscriptionService encapsulates the service that transcribes sleep talking recordings
type TranscriptionService struct {
	config *config.ServiceConfig
	reader filesystem.Reader
	speech SpeechRecogniser
}

// NewTranscriptionService initialises a Transcription Service
func NewTranscriptionService(c *config.ServiceConfig, r Reader, s SpeechRecogniser) *TranscriptionService {
	as := TranscriptionService{
		config: c,
		reader: r,
		speech: s,
	}

	return &as
}

// Start loads any recordings from the source and sends them the GCS Transcription service.Start
// The results of the request are returned on completion
func (ts TranscriptionService) Start(ctx context.Context) (TranscriptionResults, error) {
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

		resp, err := ts.speech.Recognize(ctx, &req)
		if err != nil {
			return nil, errors.Wrap(err, "error occurred recognizing speech")
		}

		results = append(results, resp)
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
