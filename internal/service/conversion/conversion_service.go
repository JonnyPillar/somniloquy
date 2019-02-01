package conversion

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/pkg/errors"
)

// Run ...
func Run(config config.ServiceConfig) (int, error) {
	files, err := ioutil.ReadDir(config.AIFFRecordingFilePath)
	if err != nil {
		return 0, errors.Wrap(err, "failed to read aiff recording dir")
	}

	var conversionCount int

	for _, f := range files {
		if !isAiffFile(f) {
			continue
		}

		fileName := f.Name()
		a := aiffFile(config, fileName)
		f := flacFile(config, fileName)

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
