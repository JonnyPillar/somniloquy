package service

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/jonnypillar/somniloquy/config"
	"github.com/pkg/errors"
)

// Reader ...
type Reader interface {
	Read() ([]os.FileInfo, error)
}

// Converter ...
type Converter interface {
	Execute(...string) error
}

// FFMPEGConverter ...
type FFMPEGConverter struct{}

// Execute ...
func (fc FFMPEGConverter) Execute(args ...string) error {
	return exec.Command("ffmpeg", args...).Run()
}

// AIFFConverter ...
type AIFFConverter struct {
	config    *config.ServiceConfig
	reader    Reader
	converter Converter
}

// NewAIFFConverter ...
func NewAIFFConverter(c *config.ServiceConfig, r Reader, conv Converter) *AIFFConverter {
	return &AIFFConverter{
		config:    c,
		reader:    r,
		converter: conv,
	}
}

// ToFlac ...
//TODO wrap the returned int with more info
func (ac AIFFConverter) ToFlac() (int, error) {
	files, err := ac.reader.Read()
	if err != nil {
		return 0, errors.Wrap(err, "failed to read aiff recording dir")
	}

	var conversionCount int

	for _, f := range files {
		if !isAiff(f) {
			continue
		}
		fileName := f.Name()
		//TODO rename methods
		a := ac.aiffFile(fileName)
		f := ac.flacFile(fileName)

		err := ac.converter.Execute("-i", a, "-c:a", "flac", f)
		if err != nil {
			return conversionCount, errors.Wrap(err, "error occured converting aiff files to flac")
		}

		fmt.Println("Converted: " + a + " to Flac:" + f)
		conversionCount++
	}

	return conversionCount, nil
}

func isAiff(f os.FileInfo) bool {
	if !f.Mode().IsRegular() {
		return false
	}

	if filepath.Ext(f.Name()) != aiffExt {
		return false
	}

	return true
}

func (ac AIFFConverter) flacFile(fileName string) string {
	flac := recordingFilePath(ac.config.FLACRecordingFilePath, fileName)

	return strings.Replace(flac, aiffExt, flacExt, 1)
}

func (ac AIFFConverter) aiffFile(fileName string) string {
	return recordingFilePath(ac.config.AIFFRecordingFilePath, fileName)
}

func recordingFilePath(filePath, fileName string) string {
	return fmt.Sprintf("%s%s", filePath, fileName)
}
