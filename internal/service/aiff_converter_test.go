package service_test

import (
	"os"
	"testing"

	"github.com/jonnypillar/somniloquy/config"
	"github.com/jonnypillar/somniloquy/internal/service"
	"github.com/jonnypillar/somniloquy/internal/service/fake"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type AiffConverterSuite struct {
	suite.Suite
}

func TestAiffConverterSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(AiffConverterSuite))
}

func (s *AiffConverterSuite) TestToFlac() {
	var tests = []struct {
		desc      string
		reader    fake.Reader
		converter fake.Converter

		expectedConvertedCount int
		expectedError          error
	}{
		{
			"given reader returns an error, error returned",
			fake.Reader{
				Files: []os.FileInfo{
					fake.FileInfo{
						FileName:  "foo.aiff",
						IsRegular: true,
					},
					fake.FileInfo{
						FileName:  "bar.aiff",
						IsRegular: true,
					},
				},
			},
			fake.Converter{},

			2,
			nil,
		},
		{
			"given reader returns an error, error returned",
			fake.Reader{
				Error: errors.New("an error occured"),
			},
			fake.Converter{},

			0,
			errors.New("failed to read aiff recording dir: an error occured"),
		},
		{
			"given the reader returns a file that isnt an .aiff file, then file is not converted",
			fake.Reader{
				Files: []os.FileInfo{
					fake.FileInfo{
						FileName:  "foo.bar",
						IsRegular: true,
					},
				},
			},
			fake.Converter{},

			0,
			nil,
		},
		{
			"given the converter returns an error, then an error is returned",
			fake.Reader{
				Files: []os.FileInfo{
					fake.FileInfo{
						FileName:  "foo.aiff",
						IsRegular: true,
					},
				},
			},
			fake.Converter{
				Error: errors.New("an error occured"),
			},

			0,
			errors.New("error occured converting aiff files to flac: an error occured"),
		},
		{
			"given the reader returns a file that is not regular, then file is not converted",
			fake.Reader{
				Files: []os.FileInfo{
					fake.FileInfo{
						FileName:  "foo.aiff",
						IsRegular: false,
					},
				},
			},
			fake.Converter{},

			0,
			nil,
		},
	}

	for _, test := range tests {
		test := test

		s.T().Run(test.desc, func(t *testing.T) {
			c := config.ServiceConfig{}
			ac := service.NewAIFFConverter(&c, &test.reader, &test.converter)

			count, err := ac.ToFlac()

			if test.expectedError != nil && s.Error(err, test.desc) {
				s.EqualError(err, test.expectedError.Error(), test.desc)
			} else {
				s.Nil(err, test.desc)
				s.Equal(test.expectedConvertedCount, count, test.desc)
			}
		})
	}
}
