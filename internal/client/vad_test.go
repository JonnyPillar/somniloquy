package client_test

import (
	"fmt"
	"testing"

	"github.com/jonnypillar/somniloquy/internal/client"
	"github.com/stretchr/testify/suite"
)

type VADSuite struct {
	suite.Suite
}

func TestVADSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(VADSuite))
}

func (s *VADSuite) TestVADFlux() {
	var tests = []struct {
		desc   string
		length int
		data   []int32

		expectedFlux float64
	}{
		{
			"given no data, a flux score of zero is returned",
			10,
			[]int32{},

			0.0,
		},
		{
			"given a set of 10 1s, a flux score of ten is returned",
			10,
			s.GenerateData(1, 10),

			10.000000000000007,
		},
		// {
		// 	"given a set of 5 1s and 5 os, a flux score of five is returned",
		// 	10,
		// 	[]int32{1, 1, 1, 1, 1, 0, 0, 0, 0, 0},

		// 	5.000000000000007,
		// },
		// {
		// 	"given a set of 5 1s and 5 os, a flux score of five is returned",
		// 	10,
		// 	[]int32{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},

		// 	5.000000000000007,
		// },
		// {
		// 	"given a set of 5 1s and 5 os, a flux score of five is returned",
		// 	10,
		// 	[]int32{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},

		// 	5.000000000000007,
		// },
		// {
		// 	"given a set of 5 1s and 5 os, a flux score of five is returned",
		// 	10,
		// 	[]int32{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},

		// 	5.000000000000007,
		// },
		// {
		// 	"given a set of 5 1s and 5 os, a flux score of five is returned",
		// 	10,
		// 	[]int32{1, 2, 3, 4, 5, 5, 4, 3, 2, 1},

		// 	5.000000000000007,
		// },
		// {
		// 	"given a set of 10 1s, a flux score of ten is returned",
		// 	10,
		// 	s.GenerateData(10, 10),

		// 	10.000000000000007,
		// },
	}

	for _, test := range tests {
		test := test

		s.T().Run(test.desc, func(t *testing.T) {
			v := client.NewVAD(test.length)

			var finalVar float64

			finalVar = v.Flux(test.data)

			fmt.Println("Testing: ", test.expectedFlux, finalVar)

			s.Equal(test.expectedFlux, finalVar, test.desc)
		})
	}
}

func (s VADSuite) GenerateData(value, times int) []int32 {
	r := []int32{}

	for i := 0; i < times; i++ {
		r = append(r, int32(value))
	}

	return r
}
