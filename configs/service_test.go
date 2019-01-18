package config_test

import (
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
		expectedConfig := config.ServiceConfig{
			Environment: "local",
			Port:        7777,
		}

		c, err := config.NewServiceConfig()

		s.Equal(c, &expectedConfig)
		s.Nil(err)
	})
}
