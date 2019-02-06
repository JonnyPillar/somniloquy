package service

import (
	"context"

	speech "cloud.google.com/go/speech/apiv1"
	"github.com/pkg/errors"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

// GoogleSpeechService ...
type GoogleSpeechService struct {
	client *speech.Client
}

// NewGoogleSpeechService ...
func NewGoogleSpeechService(ctx context.Context) (*GoogleSpeechService, error) {

	client, err := speech.NewClient(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create anew Google Cloud Speech client")
	}

	gss := GoogleSpeechService{
		client: client,
	}

	return &gss, nil
}

// Recognize ...
func (gss GoogleSpeechService) Recognize(ctx context.Context, req *speechpb.RecognizeRequest) (*TranscriptionResult, error) {
	resp, err := gss.client.Recognize(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "error occurred sending recording to Google Cloud Services API")
	}

	result := NewTranscriptionResult(resp)

	return result, nil
}
