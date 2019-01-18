package client_test

import (
	"os"

	"github.com/stretchr/testify/suite"
)

type ClientSuite struct {
	suite.Suite
}

func (cs ClientSuite) ReadFile(dir string) *os.File {
	file, err := os.Open(dir)
	if err != nil {
		panic(err)
	}

	return file
}
