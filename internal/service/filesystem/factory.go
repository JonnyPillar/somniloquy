package filesystem

import (
	"bytes"
	"os"

	"github.com/jonnypillar/somniloquy/config"
)

const (
	s3Key = "s3"
)

// Saver ...
type Saver interface {
	Save(string, *bytes.Buffer) error
}

// Reader ...
type Reader interface {
	Read() ([]os.FileInfo, error)
}

// GetReader ...
func GetReader(config *config.ServiceConfig) (Reader, error) {
	switch config.ReadDestination {
	case s3Key:
		s3, err := NewS3Bucket(config)
		if err != nil {
			return nil, err
		}

		return s3, nil
	default:
		return NewFileSystem(config), nil
	}
}

// GetSaver ...
func GetSaver(config *config.ServiceConfig) (Saver, error) {
	switch config.UploadDestination {
	case s3Key:
		s3, err := NewS3Bucket(config)
		if err != nil {
			return nil, err
		}

		return s3, nil
	default:
		return NewFileSystem(config), nil
	}
}
