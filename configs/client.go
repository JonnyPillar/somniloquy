package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

const serviceURL = "%s:%d"

// ClientConfig defines the config for the Client
type ClientConfig struct {
	ServiceHost   string  `env:"SERVICE_URL" envDefault:"localhost"`
	ServicePort   int     `env:"SERVICE_URL" envDefault:"7777"`
	SampleRate    float64 `env:"SAMPLE_RATE" envDefault:"44100"`
	SampleSeconds int     `env:"SAMPLE_SECONDS" envDefault:"7"`
}

// NewClientConfig initialises a new Client Config and sets the values based on ENV variables
func NewClientConfig() (*ClientConfig, error) {
	c := ClientConfig{}
	err := env.Parse(&c)
	if err != nil {
		return nil, errors.Wrap(err, "error occured parsing config values")
	}

	return &c, nil
}

// ServiceURL returns a formatted URL for the Services
func (c ClientConfig) ServiceURL() string {
	return fmt.Sprintf(serviceURL, c.ServiceHost, c.ServicePort)
}

// SampleDuration returns the duration that should be sampled
func (c ClientConfig) SampleDuration() time.Duration {
	return time.Duration(c.SampleSeconds) * time.Second
}
