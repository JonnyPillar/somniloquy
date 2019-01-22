package config_test

import (
	"testing"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/stretchr/testify/suite"
)

type ClientConfigSuite struct {
	suite.Suite
}

func TestClientConfigSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(ClientConfigSuite))
}

func (s *ClientConfigSuite) TestNewClientConfig() {
	s.T().Run("creates client config with default values", func(t *testing.T) {
		expectedConfig := config.ClientConfig{
			Environment:   "local",
			ServiceHost:   "localhost",
			ServicePort:   7777,
			SampleRate:    44100,
			SampleSeconds: 7,
		}

		c, err := config.NewClientConfig()

		s.Nil(err)
		s.Equal(c, &expectedConfig)
	})
}

func (s *ClientConfigSuite) TestClientConfigServiceURL() {
	var tests = []struct {
		desc string
		URL  string
		Port int

		expectedURL string
	}{
		{
			"given valid config values, fornatted service URL returned",
			"www.foo.com",
			1234,

			"www.foo.com:1234",
		},
	}

	for _, test := range tests {
		test := test

		s.T().Run(test.desc, func(t *testing.T) {
			c := config.ClientConfig{
				ServiceHost: test.URL,
				ServicePort: test.Port,
			}

			url := c.ServiceURL()

			s.Equal(test.expectedURL, url, test.desc)
		})
	}
}

func (s *ClientConfigSuite) TestClientConfigSampleDuration() {
	var tests = []struct {
		desc          string
		sampleSeconds int

		expecteDuration float64
	}{
		{
			"given valid config values, correct duration returned",
			4,

			4,
		},
	}

	for _, test := range tests {
		test := test

		s.T().Run(test.desc, func(t *testing.T) {
			c := config.ClientConfig{
				SampleSeconds: test.sampleSeconds,
			}

			d := c.SampleDuration()

			s.Equal(test.expecteDuration, d.Seconds(), test.desc)
		})
	}
}
