package config

import (
	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

// ServiceConfig defines the config for the Services
type ServiceConfig struct {
	Environment string `env:"ENV" envDefault:"local"`
	Port        int    `env:"PORT" envDefault:"7777"`
}

// NewServiceConfig initialises a new Service Config and sets the values based on ENV variables
func NewServiceConfig() (*ServiceConfig, error) {
	c := ServiceConfig{}
	err := env.Parse(&c)
	if err != nil {
		return nil, errors.Wrap(err, "error occured creating Client Config")
	}

	return &c, nil
}
