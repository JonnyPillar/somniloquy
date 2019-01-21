package client_test

import (
	"errors"
	"testing"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/jonnypillar/somniloquy/internal/client"

	"github.com/stretchr/testify/suite"
)

type MicrophoneSuite struct {
	suite.Suite
}

func TestMicrophoneSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(MicrophoneSuite))
}
func (s *MicrophoneSuite) TestNewMicrophone() {
	var tests = []struct {
		desc   string
		config config.ClientConfig

		expectedError error
	}{
		{
			"given no errors occur, initialised Microphone is returned",
			config.ClientConfig{
				SampleRate: 44100,
			},

			nil,
		},
		{
			"given an error occurred while creating a Microphone, error is returned",
			config.ClientConfig{},

			errors.New("error occurred creating new microphone input: Invalid sample rate"),
		},
	}

	for _, test := range tests {
		test := test

		s.T().Run(test.desc, func(t *testing.T) {
			m, err := client.NewMicrophone(&test.config)

			if test.expectedError != nil && s.Error(err) {
				s.EqualError(err, test.expectedError.Error(), test.desc)
			} else {
				s.Nil(err, test.desc)
				s.NotNil(m, test.desc)
			}
		})
	}
}
