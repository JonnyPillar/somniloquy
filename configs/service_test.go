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
			UploadDestinations:      []string{"file"},
			AWSUploadS3BucketName:   "dev-somniloquy-uploads",
			AWSRegion:               "eu-west-1",
		}

		c, err := config.NewServiceConfig()

		s.Equal(c, &expectedConfig)
		s.Nil(err)
	})
}
