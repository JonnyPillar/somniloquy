package filesystem_test

import (
	"testing"

	"github.com/jonnypillar/somniloquy/config"
	"github.com/jonnypillar/somniloquy/internal/service/filesystem"
	"github.com/stretchr/testify/suite"
)

type FactorySuite struct {
	suite.Suite
}

func TestFactorySuiteTestSuite(t *testing.T) {
	suite.Run(t, new(FactorySuite))
}

func (s *FactorySuite) TestGetReader() {
	var tests = []struct {
		desc              string
		readerDestination string

		expectedReader filesystem.Reader
		expectedError  error
	}{
		{
			"given file destination, file system reader returned",
			"file",

			&filesystem.FileSystem{},
			nil,
		},
	}

	for _, test := range tests {
		test := test

		s.T().Run(test.desc, func(t *testing.T) {
			c := config.ServiceConfig{
				ReadDestination: test.readerDestination,
			}

			r, err := filesystem.GetReader(&c)

			if test.expectedError != nil && s.Error(err) {
				s.EqualError(err, test.expectedError.Error(), test.desc)
			} else {
				s.Nil(err, test.desc)
				s.IsType(test.expectedReader, r, test.desc)
			}
		})
	}
}

func (s *FactorySuite) TestGetSaver() {
	var tests = []struct {
		desc              string
		readerDestination string

		expectedSaver filesystem.Saver
		expectedError error
	}{
		{
			"given file destination, file system saver returned",
			"file",

			&filesystem.FileSystem{},
			nil,
		},
	}

	for _, test := range tests {
		test := test

		s.T().Run(test.desc, func(t *testing.T) {
			c := config.ServiceConfig{
				ReadDestination: test.readerDestination,
			}

			r, err := filesystem.GetSaver(&c)

			if test.expectedError != nil && s.Error(err) {
				s.EqualError(err, test.expectedError.Error(), test.desc)
			} else {
				s.Nil(err, test.desc)
				s.IsType(test.expectedSaver, r, test.desc)
			}
		})
	}
}
