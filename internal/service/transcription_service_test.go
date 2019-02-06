package service_test

import (
	"context"
	"os"
	"testing"

	"github.com/jonnypillar/somniloquy/config"
	"github.com/jonnypillar/somniloquy/internal/service"
	"github.com/jonnypillar/somniloquy/internal/service/fake"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type TranscriptionServiceSuite struct {
	suite.Suite
}

func TestTranscriptionServiceSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(TranscriptionServiceSuite))
}
func (s *TranscriptionServiceSuite) TestStart() {
	var tests = []struct {
		desc   string
		reader fake.Reader
		speech fake.SpeechRecogniser

		expectedResults service.TranscriptionResults
		expectedError   error
	}{
		{
			"given the reader returns an error, an error is returned",
			fake.Reader{
				Error: errors.New("an error occured"),
			},
			fake.SpeechRecogniser{},

			nil,
			errors.New("failed to read flac recording dir: an error occured"),
		},
		{
			"given the reader returns a file that is not a .flac file, file is not transcribed",
			fake.Reader{
				Files: []os.FileInfo{
					fake.FileInfo{
						FileName:  "foo.bar",
						IsRegular: true,
					},
				},
			},
			fake.SpeechRecogniser{},

			service.TranscriptionResults{},
			nil,
		},
	}

	for _, test := range tests {
		test := test

		s.T().Run(test.desc, func(t *testing.T) {
			c := config.ServiceConfig{}
			ctx := context.Background()

			ts := service.NewTranscriptionService(&c, test.reader, test.speech)

			_, err := ts.Start(ctx)

			if test.expectedError != nil && s.Error(err, test.desc) {
				s.EqualError(err, test.expectedError.Error(), test.desc)
			} else {
				s.Nil(err, test.desc)
			}
		})
	}
}
