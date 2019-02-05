package record_test

import (
	"testing"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/jonnypillar/somniloquy/internal/api"
	"github.com/jonnypillar/somniloquy/internal/service/record"
	"github.com/jonnypillar/somniloquy/internal/service/record/fake"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type RecordingServiceSuite struct {
	suite.Suite
}

func TestRecordingServiceSuiteTestSuite(t *testing.T) {
	suite.Run(t, new(RecordingServiceSuite))
}

func (s *RecordingServiceSuite) TestUpload() {
	var tests = []struct {
		desc   string
		stream fake.UploadStream
		saver  fake.Saver

		expectedStatus api.UploadStatus
		expectedError  error
	}{
		{
			desc: "given stream data is saved successfully, no error is returned & OK is returned to stream",
			stream: fake.UploadStream{
				Content: []int32{1, 2, 3, 4, 5},
			},
			saver: fake.Saver{},

			expectedStatus: api.UploadStatus{
				Message: "Upload received with success",
				Code:    api.UploadStatusCode_Ok,
			},
		},
		{
			desc: "given no stream content, error is returned",
			stream: fake.UploadStream{
				Content: []int32{},
			},

			expectedError: errors.New("failed to create recording buffer: no data to encode"),
		},
		{
			desc: "given an error occurs receiving data from stream, error is returned",
			stream: fake.UploadStream{
				RecvError: errors.New("error occured"),
			},

			expectedError: errors.New("failed unexpectadely while reading chunks from stream: error occured"),
		},
		{
			desc: "given an error occurs receiving data from stream, error is returned",
			stream: fake.UploadStream{
				Content: []int32{1, 2, 3, 4, 5},
			},
			saver: fake.Saver{
				Error: errors.New("error occured"),
			},

			expectedError: errors.New("error occured saving recording: error occured"),
		},
		{
			desc: "given an error occurs closing stream, error is returned",
			stream: fake.UploadStream{
				Content:             []int32{1, 2, 3, 4, 5},
				StreamAndCloseError: errors.New("error occured"),
			},
			saver: fake.Saver{},

			expectedError: errors.New("failed to send status code: error occured"),
		},
	}

	for _, test := range tests {
		test := test

		s.T().Run(test.desc, func(t *testing.T) {
			c := config.ServiceConfig{}
			rs := record.NewRecordingService(&c, test.saver)

			//Subject
			err := rs.Upload(&test.stream)

			if test.expectedError != nil && s.Error(err) {
				s.EqualError(err, test.expectedError.Error(), test.desc)
			} else {
				s.Nil(err, test.desc)
				s.Equal(&test.expectedStatus, test.stream.ReceviedStatus, test.desc)
			}
		})
	}
}
