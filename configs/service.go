package config

import (
	"github.com/caarlos0/env"
	"github.com/pkg/errors"
)

// ServiceConfig defines the config for the Services
type ServiceConfig struct {
	Port                    int      `env:"PORT" envDefault:"7777"`
	GoogleAppServicesConfig string   `env:"GOOGLE_APPLICATION_CREDENTIALS" envDefault:"~/.gcs/config.json"`
	SampleRate              int      `env:"SAMPLE_RATE" envDefault:"44100"`
	UploadDestinations      []string `env:"UPLOAD_DESTINATIONS" envDefault:"file"`
	RecordingFilePath       string   `env:"RECORDING_FILE_PATH" envDefault:"./assets/recordings/aiff/"`
	AWSRegion               string   `env:"AWS_REGION" envDefault:"eu-west-1"`
	AWSUploadS3BucketName   string   `env:"AWS_UPLOAD_S3_BUCKET_NAME" envDefault:"dev-somniloquy-uploads"`
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
