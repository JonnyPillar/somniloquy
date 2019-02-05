package fake

import (
	"context"
	"io"

	"github.com/jonnypillar/somniloquy/internal/api"
	"google.golang.org/grpc/metadata"
)

// UploadStream ...
type UploadStream struct {
	Content             []int32
	RecvError           error
	StreamAndCloseError error

	ReceviedStatus *api.UploadStatus
}

// SendAndClose ...
func (us *UploadStream) SendAndClose(s *api.UploadStatus) error {
	us.ReceviedStatus = s

	if us.StreamAndCloseError != nil {
		return us.StreamAndCloseError
	}

	return nil
}

// Recv ...
func (us *UploadStream) Recv() (*api.UploadRecordRequest, error) {
	if us.RecvError != nil {
		return nil, us.RecvError
	}

	if len(us.Content) == 0 {
		return nil, io.EOF
	}

	r := &api.UploadRecordRequest{
		Content: us.Content,
	}

	us.Content = us.Content[:0]

	return r, nil
}

// SetHeader ...
func (us UploadStream) SetHeader(metadata.MD) error {
	panic("not implemented")
}

// SendHeader ...
func (us UploadStream) SendHeader(metadata.MD) error {
	panic("not implemented")
}

// SetTrailer ...
func (us UploadStream) SetTrailer(metadata.MD) {
	panic("not implemented")
}

// Context ...
func (us UploadStream) Context() context.Context {
	panic("not implemented")
}

// SendMsg ...
func (us UploadStream) SendMsg(m interface{}) error {
	panic("not implemented")
}

// RecvMsg ...
func (us UploadStream) RecvMsg(m interface{}) error {
	panic("not implemented")
}
