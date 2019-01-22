package config_test

import (
	"os"
	"testing"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/stretchr/testify/suite"
)

type ServiceConfigSuite struct {
	suite.Suite
}

func TestServiceConfigSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceConfigSuite))
}

func (s *ServiceConfigSuite) TestNewServiceConfig() {
	s.T().Run("creates client config with default values", func(t *testing.T) {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "~/.gcs/config.json")

		expectedConfig := config.ServiceConfig{
			Environment:             "local",
			Port:                    7777,
			GoogleAppServicesConfig: "~/.gcs/config.json",
			SampleRate:              44100,
		}

		c, err := config.NewServiceConfig()

		s.Nil(err)
		s.Equal(c, &expectedConfig)
	})
}
