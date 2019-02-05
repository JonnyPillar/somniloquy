package service

import (
	"fmt"
	"strings"

	"google.golang.org/genproto/googleapis/cloud/speech/v1"
)

// TranscriptionResults defines a set of GCS Transcription Results
type TranscriptionResults []*TranscriptionResult

// TranscriptionResult defines results for a set of GCS Transcription requests
type TranscriptionResult struct {
	Results []Transcription
}

// Transcription defines a GSC Transcription result
type Transcription struct {
	Transcription string
	Confidence    float32
}

// NewTranscriptionResult creates a TranscriptionResult from a raw GCS Speech RecognizeResponse
// If a GCS Transcription Result contains more than one alternative, we only capture the first Alternative
func NewTranscriptionResult(resp *speech.RecognizeResponse) *TranscriptionResult {
	c := TranscriptionResult{}

	for _, res := range resp.Results {
		if len(res.Alternatives) == 0 {
			continue
		}

		alt := res.Alternatives[0]

		c.Results = append(c.Results, Transcription{
			Transcription: alt.Transcript,
			Confidence:    alt.Confidence,
		})
	}

	return &c
}

// String ...
func (tr TranscriptionResults) String() string {
	b := strings.Builder{}

	b.WriteString("Transcription Results\n\n")

	for _, r := range tr {
		b.WriteString(r.String())
	}

	return b.String()
}

// String ...
func (tr TranscriptionResult) String() string {
	b := strings.Builder{}

	for _, r := range tr.Results {
		b.WriteString("- Confidence: ")
		b.WriteString(fmt.Sprintf("%f", r.Confidence))
		b.WriteString("\n- Transcription: ")
		b.WriteString(r.Transcription)
	}

	b.WriteString("\n")

	return b.String()
}
