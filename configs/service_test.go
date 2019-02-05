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
			Port:                    7777,
			GoogleAppServicesConfig: "~/.gcs/config.json",
			SampleRate:              44100,
			UploadDestination:       "file",
			ReadDestination:         "file",
			AIFFRecordingFilePath:   "./assets/recordings/aiff/",
			FLACRecordingFilePath:   "./assets/recordings/flac/",
			AWSRegion:               "eu-west-1",
			AWSUploadS3BucketName:   "dev-somniloquy-uploads",
		}

		c, err := config.NewServiceConfig()

		s.Equal(c, &expectedConfig)
		s.Nil(err)
	})
}
