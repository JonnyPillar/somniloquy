package service

import (
	"context"

	"cloud.google.com/go/speech/apiv1"
	"github.com/jonnypillar/somniloquy/configs"
	"github.com/pkg/errors"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

const languageCode = "en-GB"

// TranscriptionService ...
type TranscriptionService struct {
	ctx    context.Context
	config *config.ServiceConfig
	client *speech.Client
}

// TranscriptionResponse ...
type TranscriptionResponse struct {
	results []TranscriptionResponseResult
}

// TranscriptionResponseResult ...
type TranscriptionResponseResult struct {
	transcription string
	confidence    float32
}

// NewTranscriptionService ...
func NewTranscriptionService(ctx context.Context, config *config.ServiceConfig, client *speech.Client) *TranscriptionService {
	return &TranscriptionService{
		ctx:    ctx,
		config: config,
		client: client,
	}
}

// NewTranscriptionResponse ...
func NewTranscriptionResponse(resp *speechpb.RecognizeResponse) *TranscriptionResponse {
	c := TranscriptionResponse{}

	for _, res := range resp.Results {
		if len(res.Alternatives) == 0 {
			return &c
		}

		alt := res.Alternatives[0]

		c.results = append(c.results, TranscriptionResponseResult{
			transcription: alt.Transcript,
			confidence:    alt.Confidence,
		})
	}

	return &c
}

// Run ...
func (ts TranscriptionService) Run(data []byte) (*TranscriptionResponse, error) {
	config := speechpb.RecognitionConfig{
		Encoding:        speechpb.RecognitionConfig_FLAC,
		SampleRateHertz: int32(ts.config.SampleRate),
		LanguageCode:    languageCode,
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

	resp, err := ts.client.Recognize(ts.ctx, &req)
	if err != nil {
		return nil, errors.Wrap(err, "error occured trying to recognize speech")
	}

	res := NewTranscriptionResponse(resp)

	return res, nil
}
