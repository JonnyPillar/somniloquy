package service_test

import (
	"testing"

	"github.com/jonnypillar/somniloquy/internal/service"
	"github.com/stretchr/testify/suite"
	"google.golang.org/genproto/googleapis/cloud/speech/v1"
)

type TranscriptionResultsSuite struct {
	suite.Suite
}

func TestTranscriptionResultsSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(TranscriptionResultsSuite))
}

func (s *TranscriptionResultsSuite) TestNewTranscriptionResult() {
	var tests = []struct {
		desc     string
		response speech.RecognizeResponse

		expectedResult service.TranscriptionResult
	}{
		{
			"given an empty set of response results, empty transcription result returned",
			speech.RecognizeResponse{
				Results: []*speech.SpeechRecognitionResult{},
			},

			service.TranscriptionResult{},
		},
		{
			"given set of response results, transcription result returned",
			speech.RecognizeResponse{
				Results: []*speech.SpeechRecognitionResult{
					&speech.SpeechRecognitionResult{
						Alternatives: []*speech.SpeechRecognitionAlternative{
							&speech.SpeechRecognitionAlternative{
								Transcript: "Hello World",
								Confidence: 12.2,
							},
						},
					},
					&speech.SpeechRecognitionResult{
						Alternatives: []*speech.SpeechRecognitionAlternative{
							&speech.SpeechRecognitionAlternative{
								Transcript: "Good Bye",
								Confidence: 88.2,
							},
						},
					},
				},
			},

			service.TranscriptionResult{
				Results: []service.Transcription{
					service.Transcription{
						Transcription: "Hello World",
						Confidence:    12.2,
					},
					service.Transcription{
						Transcription: "Good Bye",
						Confidence:    88.2,
					},
				},
			},
		},
		{
			"given there are more than one alternatives in a result, transcription result containing only the first alternative returned",
			speech.RecognizeResponse{
				Results: []*speech.SpeechRecognitionResult{
					&speech.SpeechRecognitionResult{
						Alternatives: []*speech.SpeechRecognitionAlternative{
							&speech.SpeechRecognitionAlternative{
								Transcript: "Hello World",
								Confidence: 12.2,
							},
							&speech.SpeechRecognitionAlternative{
								Transcript: "World Hello",
								Confidence: 1.2,
							},
						},
					},
				},
			},

			service.TranscriptionResult{
				Results: []service.Transcription{
					service.Transcription{
						Transcription: "Hello World",
						Confidence:    12.2,
					},
				},
			},
		},
		{
			"given a result has no alternatives, result is skipped",
			speech.RecognizeResponse{
				Results: []*speech.SpeechRecognitionResult{
					&speech.SpeechRecognitionResult{
						Alternatives: []*speech.SpeechRecognitionAlternative{},
					},
				},
			},

			service.TranscriptionResult{},
		},
	}

	for _, test := range tests {
		test := test

		s.T().Run(test.desc, func(t *testing.T) {
			tr := service.NewTranscriptionResult(&test.response)

			s.Equal(&test.expectedResult, tr, test.desc)
		})
	}
}

func (s *TranscriptionResultsSuite) TestTranscriptionResultsString() {
	var tests = []struct {
		desc                 string
		transcriptionResults service.TranscriptionResults

		expectedString string
	}{
		{
			"given there are no transcription results",
			service.TranscriptionResults{},

			"Transcription Results\n\n",
		},
		{
			"given transcription results",
			service.TranscriptionResults{
				&service.TranscriptionResult{
					Results: []service.Transcription{
						service.Transcription{
							Transcription: "Hello World",
							Confidence:    12.2,
						},
					},
				},
			},

			"Transcription Results\n\n- Confidence: 12.200000\n- Transcription: Hello World\n",
		},
	}

	for _, test := range tests {
		test := test

		s.T().Run(test.desc, func(t *testing.T) {
			res := test.transcriptionResults.String()

			s.Equal(test.expectedString, res, test.desc)
		})
	}
}
