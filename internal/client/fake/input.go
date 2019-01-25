package fake

import "github.com/jonnypillar/somniloquy/internal/api"

// Input ...
type Input struct {
	SendError    error
	UploadStatus api.UploadStatus
	UploadError  error
}

// Send ...
func (i Input) Send(*api.UploadRecordRequest) error {
	return i.SendError
}

// CloseAndRecv ...
func (i Input) CloseAndRecv() (*api.UploadStatus, error) {
	return &i.UploadStatus, i.UploadError
}
