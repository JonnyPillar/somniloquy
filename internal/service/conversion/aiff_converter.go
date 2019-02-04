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
	//TODO probably not a os.FileInfo from S3
	Read(string) ([]os.FileInfo, error)
	Exists(string) (bool, error)
}

// AIFFConverter ...
type AIFFConverter struct {
	config config.ServiceConfig
	reader Reader
}

// NewAIFFConverter ...
func NewAIFFConverter(config config.ServiceConfig, reader Reader) *AIFFConverter {
	return &AIFFConverter{
		config: config,
		reader: reader,
	}
}

// ToFlac ...
func (ac AIFFConverter) ToFlac() (int, error) {
	files, err := ac.reader.Read(ac.config.AIFFRecordingFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "failed to read aiff recording dir")
	}

	var conversionCount int

	for _, f := range files {
		if !isAiffFile(f) {
			continue
		}

		fileName := f.Name()
		//TODO as this is no longer happening on the local file system we will need to create a temp AIFF file and convert it to FLAC.
		//This could happen at the recording stage??
		a := aiffFile(ac.config, fileName)
		f := flacFile(ac.config, fileName)

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

func flacFile(config config.ServiceConfig, fileName string) string {
	flac := recordingFilePath(config.FLACRecordingFilePath, fileName)

	return strings.Replace(flac, ".aiff", ".flac", 1)
}

func aiffFile(config config.ServiceConfig, fileName string) string {
	return recordingFilePath(config.AIFFRecordingFilePath, fileName)
}

func recordingFilePath(filePath, fileName string) string {
	return fmt.Sprintf("%s%s", filePath, fileName)
}
