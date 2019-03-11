package client_test

// import (
// 	"context"
// 	"errors"
// 	"testing"

// 	"github.com/jonnypillar/somniloquy/config"
// 	"github.com/jonnypillar/somniloquy/internal/api"
// 	"github.com/jonnypillar/somniloquy/internal/client"
// 	"github.com/jonnypillar/somniloquy/internal/client/fake"
// 	"github.com/stretchr/testify/suite"
// 	"google.golang.org/grpc"
// )

// type ClientSuite struct {
// 	suite.Suite
// }

// func TestClientSuiteTestSuite(t *testing.T) {
// 	suite.Run(t, new(ClientSuite))
// }

// func (s *ClientSuite) TestNewClient() {
// 	var tests = []struct {
// 		desc   string
// 		config config.ClientConfig

// 		expectedError error
// 	}{
// 		{
// 			"given no errors occur, initialised Client is returned",
// 			config.ClientConfig{
// 				SampleRate: 44100,
// 			},

// 			nil,
// 		},
// 		{
// 			"given an error occurred while creating a Client, error is returned",
// 			config.ClientConfig{},

// 			errors.New("error occurred creating new Client: error occurred creating new microphone input: Invalid sample rate"),
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test

// 		s.T().Run(test.desc, func(t *testing.T) {
// 			m, err := client.NewClient(&test.config, &grpc.ClientConn{})

// 			if test.expectedError != nil && s.Error(err) {
// 				s.EqualError(err, test.expectedError.Error(), test.desc)
// 			} else {
// 				s.Nil(err, test.desc)
// 				s.NotNil(m, test.desc)
// 			}
// 		})
// 	}
// }

// func (s *ClientSuite) TestClientSend() {
// 	var tests = []struct {
// 		desc   string
// 		config config.ClientConfig
// 		input  fake.Input

// 		expectedError error
// 	}{
// 		{
// 			"given client successfully streams, nil is returned",
// 			config.ClientConfig{
// 				SampleRate:    44100,
// 				SampleSeconds: 1,
// 			},
// 			fake.Input{
// 				UploadStatus: api.UploadStatus{
// 					Code: api.UploadStatusCode_Ok,
// 				},
// 			},

// 			nil,
// 		},
// 		{
// 			"given an error occurs sending stream, error is returned",
// 			config.ClientConfig{
// 				SampleRate:    44100,
// 				SampleSeconds: 1,
// 			},
// 			fake.Input{
// 				SendError: errors.New("an error occured"),
// 			},

// 			errors.New("error occured recording microphone: error occured sending chunk: an error occured"),
// 		},
// 		{
// 			"given an error occurs closing and receiving stream, error is returned",
// 			config.ClientConfig{
// 				SampleRate:    44100,
// 				SampleSeconds: 1,
// 			},
// 			fake.Input{
// 				UploadError: errors.New("an error occured"),
// 			},

// 			errors.New("failed to receive upstream status response: an error occured"),
// 		},
// 		{
// 			"given a failed status code is returned from the service, error is returned",
// 			config.ClientConfig{
// 				SampleRate:    44100,
// 				SampleSeconds: 1,
// 			},
// 			fake.Input{
// 				UploadStatus: api.UploadStatus{
// 					Code:    api.UploadStatusCode_Failed,
// 					Message: "an error occured",
// 				},
// 			},

// 			errors.New("failed to upload stream. an error occured"),
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test

// 		s.T().Run(test.desc, func(t *testing.T) {
// 			c, _ := client.NewClient(&test.config, &grpc.ClientConn{})

// 			err := c.Stream(context.Background(), test.input)

// 			if test.expectedError != nil && s.Error(err) {
// 				s.EqualError(err, test.expectedError.Error(), test.desc)
// 			} else {
// 				s.Nil(err, test.desc)
// 			}
// 		})
// 	}
// }
