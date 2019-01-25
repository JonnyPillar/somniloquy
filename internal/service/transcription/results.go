package transcription

import (
	"fmt"
	"strings"

	"google.golang.org/genproto/googleapis/cloud/speech/v1"
)

// Results ...
type Results []*GCSResults

// GCSResults ...
type GCSResults struct {
	results []GCSResult
}

// GCSResult ...
type GCSResult struct {
	transcription string
	confidence    float32
}

func newGCSResults(resp *speech.RecognizeResponse) *GCSResults {
	c := GCSResults{}

	for _, res := range resp.Results {
		if len(res.Alternatives) == 0 {
			return &c
		}

		alt := res.Alternatives[0]

		c.results = append(c.results, GCSResult{
			transcription: alt.Transcript,
			confidence:    alt.Confidence,
		})
	}

	return &c
}

func (tr Results) String() string {
	b := strings.Builder{}

	b.WriteString("Transcription Results\n\n")

	for _, r := range tr {
		b.WriteString(r.String())
	}

	return b.String()
}

func (tr GCSResults) String() string {
	b := strings.Builder{}

	for _, r := range tr.results {
		b.WriteString("- Confidence: ")
		b.WriteString(fmt.Sprintf("%f", r.confidence))
		b.WriteString("\n- Transcription: ")
		b.WriteString(r.transcription)
	}

	b.WriteString("\n")

	return b.String()
}
