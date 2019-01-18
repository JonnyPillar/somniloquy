package config

import (
	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

// ServiceConfig ...
type ServiceConfig struct {
	Environment string `env:"ENV" envDefault:"local"`
	Port        int    `env:"PORT" envDefault:"7777"`
}

// NewServiceConfig ...
func NewServiceConfig() (*ServiceConfig, error) {
	c := ServiceConfig{}
	err := env.Parse(&c)
	if err != nil {
		return nil, errors.Wrap(err, "error occured creating Client Config")
	}

	return &c, nil
}
