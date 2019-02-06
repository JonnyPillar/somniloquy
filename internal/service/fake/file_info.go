package fake

import (
	"os"
	"time"
)

// FileInfo ...
type FileInfo struct {
	FileName  string
	IsRegular bool
}

// Name ...
func (f FileInfo) Name() string {
	return f.FileName
}

// Size ...
func (f FileInfo) Size() int64 {
	panic("not implemented")
}

// Mode ...
func (f FileInfo) Mode() os.FileMode {
	if !f.IsRegular {
		return os.ModeIrregular
	}

	return 0
}

// ModTime ...
func (f FileInfo) ModTime() time.Time {
	panic("not implemented")
}

// IsDir ...
func (f FileInfo) IsDir() bool {
	panic("not implemented")
}

// Sys ...
func (f FileInfo) Sys() interface{} {
	panic("not implemented")
}
