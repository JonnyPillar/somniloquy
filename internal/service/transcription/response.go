package transcription

import (
	"fmt"
	"strings"

	"google.golang.org/genproto/googleapis/cloud/speech/v1"
)

type transcriptionResponse struct {
	results []transcriptionResponseResult
}

type transcriptionResponseResult struct {
	transcription string
	confidence    float32
}

func newTranscriptionResponse(resp *speech.RecognizeResponse) *transcriptionResponse {
	c := transcriptionResponse{}

	for _, res := range resp.Results {
		if len(res.Alternatives) == 0 {
			return &c
		}

		alt := res.Alternatives[0]

		c.results = append(c.results, transcriptionResponseResult{
			transcription: alt.Transcript,
			confidence:    alt.Confidence,
		})
	}

	return &c
}

func (tr *transcriptionResponse) String() string {
	b := strings.Builder{}

	for _, r := range tr.results {
		b.WriteString(r.String())
	}

	return b.String()
}

func (tr *transcriptionResponseResult) String() string {
	return fmt.Sprintf("Transcription: %s\nConfidence: %f", tr.transcription, tr.confidence)
}
