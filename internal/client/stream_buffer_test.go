package client_test

import (
	"errors"
	"io"
	"testing"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/jonnypillar/somniloquy/internal/client"
	"github.com/stretchr/testify/suite"
)

type StreamBufferSuite struct {
	ClientSuite
}

func TestStreamBufferSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(StreamBufferSuite))
}

func (s *StreamBufferSuite) TestStreamBufferRecord() {

	var tests = []struct {
		desc   string
		source io.Reader

		expectedChunkCount int
		expectedError      error
	}{
		{
			"given a test audio file is used as the source, correct number of chunks returned through the channel",
			s.ReadFile("../../test/data/test.mp3"),

			187,
			nil,
		},
		{
			"given an error occurs reading from the source, error is returned",
			FakeFileSource{
				err: errors.New("error occured"),
			},

			0,
			errors.New("error occurred reading byte array: error occured"),
		},
	}

	for _, test := range tests {
		test := test

		s.T().Run(test.desc, func(t *testing.T) {
			config, _ := config.NewClientConfig()
			chunks, err := client.BufferStream(config, test.source)

			if test.expectedError != nil && s.Error(err) {
				s.EqualError(err, test.expectedError.Error(), test.desc)
			} else {
				s.Nil(err, test.desc)
				s.Equal(test.expectedChunkCount, len(chunks))
			}
		})
	}
}

// FakeFileSource ...
type FakeFileSource struct {
	err error
}

// Read ...
func (fs FakeFileSource) Read(p []byte) (n int, err error) {
	return 0, fs.err
}

// Close ...
func (fs FakeFileSource) Close() error {
	return fs.err
}
