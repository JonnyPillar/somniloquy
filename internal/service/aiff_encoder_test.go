package service_test

import (
	"testing"

	"github.com/jonnypillar/somniloquy/internal/service"
	"github.com/stretchr/testify/suite"
)

type AiffEncoderSuite struct {
	suite.Suite
}

func TestAiffEncoderSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(AiffEncoderSuite))
}

func (s *AiffEncoderSuite) TestAiffEncoderAppend() {
	var tests = []struct {
		desc string
		data []int32

		expectedNumberOfSamples int
	}{
		{
			"given a data of length 5, sample size is 5",
			[]int32{1, 2, 3, 4, 5},

			5,
		},
	}

	for _, test := range tests {
		test := test

		s.T().Run(test.desc, func(t *testing.T) {
			ae := service.NewAiffEncoder()

			ae.Append(test.data)

			s.Equal(test.expectedNumberOfSamples, ae.SampleCount(), test.desc)
		})
	}
}
