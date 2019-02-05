package conversion

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/pkg/errors"
)

// Reader ...
type Reader interface {
	Read() ([]os.FileInfo, error)
}

// AIFFConverter ...
type AIFFConverter struct {
	config *config.ServiceConfig
	reader Reader
}

// NewAIFFConverter ...
func NewAIFFConverter(config *config.ServiceConfig, reader Reader) *AIFFConverter {
	return &AIFFConverter{
		config: config,
		reader: reader,
	}
}

// ToFlac ...
func (ac AIFFConverter) ToFlac() (int, error) {
	files, err := ac.reader.Read()
	if err != nil {
		return 0, errors.Wrap(err, "failed to read aiff recording dir")
	}

	var conversionCount int

	for _, f := range files {
		if !isAiffFile(f) {
			continue
		}

		fileName := f.Name()
		//TODO rename methods
		a := ac.aiffFile(fileName)
		f := ac.flacFile(fileName)

		if flacExists(f) {
			continue
		}

		err := exec.Command("ffmpeg", "-i", a, "-c:a", "flac", f).Run()
		if err != nil {
			return conversionCount, errors.Wrap(err, "error occured converting aiff files to flac")
		}

		fmt.Println("Converted: " + a + " to Flac:" + f)
		conversionCount++
	}

	return conversionCount, nil
}

func isAiffFile(f os.FileInfo) bool {
	if !f.Mode().IsRegular() {
		return false
	}

	if filepath.Ext(f.Name()) != ".aiff" {
		return false
	}

	return true
}

func flacExists(filename string) bool {
	_, err := os.Stat(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func (ac AIFFConverter) flacFile(fileName string) string {
	flac := recordingFilePath(ac.config.FLACRecordingFilePath, fileName)

	return strings.Replace(flac, ".aiff", ".flac", 1)
}

func (ac AIFFConverter) aiffFile(fileName string) string {
	return recordingFilePath(ac.config.AIFFRecordingFilePath, fileName)
}

func recordingFilePath(filePath, fileName string) string {
	return fmt.Sprintf("%s%s", filePath, fileName)
}
