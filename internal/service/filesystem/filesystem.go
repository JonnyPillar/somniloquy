package filesystem

import (
	"bytes"
	"fmt"
	"os"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/pkg/errors"
)

// FileSystem ...
//TODO rename this or the package name
type FileSystem struct {
	dir string
}

// NewFileSystem ...
func NewFileSystem(config *config.ServiceConfig) *FileSystem {
	fmt.Println("Initilising FileSystem Saver")
	return &FileSystem{
		dir: config.AIFFRecordingFilePath,
	}
}

// Save ...
func (f FileSystem) Save(filename string, by *bytes.Buffer) error {
	filePath := f.dir + filename
	file, err := os.Create(filePath)
	if err != nil {
		return errors.Wrap(err, "error occured saving to file system")
	}
	defer file.Close()

	_, err = file.Write(by.Bytes())
	if err != nil {
		return err
	}

	fmt.Println("Saved file to file system:", filePath)
	return nil
}

// Read ...
func (f FileSystem) Read() ([]os.FileInfo, error) {
	return nil, nil
}
