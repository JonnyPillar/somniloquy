package fake

import "os"

// Reader ...
type Reader struct {
	Files []os.FileInfo
	Error error
}

// Save ...
func (r Reader) Read() ([]os.FileInfo, error) {
	if r.Error != nil {
		return nil, r.Error
	}

	return r.Files, nil
}
