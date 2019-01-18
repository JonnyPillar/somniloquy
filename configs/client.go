package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

const serviceURL = "%s:%d"

// ClientConfig ...
type ClientConfig struct {
	Environment string `env:"ENV" envDefault:"local"`
	ServiceHost string `env:"SERVICE_URL" envDefault:"localhost"`
	ServicePort int    `env:"SERVICE_URL" envDefault:"7777"`
}

// NewClientConfig ...
func NewClientConfig() (*ClientConfig, error) {
	c := ClientConfig{}
	err := env.Parse(&c)
	if err != nil {
		return nil, errors.Wrap(err, "error occured parsing config values")
	}

	return &c, nil
}

// ServiceURL ...
func (c ClientConfig) ServiceURL() string {
	return fmt.Sprintf(serviceURL, c.ServiceHost, c.ServicePort)
}
