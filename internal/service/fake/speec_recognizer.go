package fake

import (
	"context"

	"github.com/jonnypillar/somniloquy/internal/service"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

// SpeechRecogniser ...
type SpeechRecogniser struct {
	Error error
}

// Recognize ...
func (sr SpeechRecogniser) Recognize(ctx context.Context, req *speechpb.RecognizeRequest) (*service.TranscriptionResult, error) {
	if sr.Error != nil {
		return nil, sr.Error
	}

	return nil, nil
}
