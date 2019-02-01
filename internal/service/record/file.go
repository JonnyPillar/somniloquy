package record

import (
	"bytes"
	"fmt"
	"os"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/pkg/errors"
)

// File ...
type File struct {
	dir string
}

// NewFile ...
func NewFile(config *config.ServiceConfig) *File {
	fmt.Println("Initilising File Saver")
	return &File{
		dir: config.RecordingFilePath,
	}
}

// Save ...
func (f File) Save(filename string, by *bytes.Buffer) error {
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
