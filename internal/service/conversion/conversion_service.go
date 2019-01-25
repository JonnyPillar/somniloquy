package conversion

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

const (
	aiffDir = "./assets/recordings/aiff/"
	flacDir = "./assets/recordings/flac/"
)

// Run ...
func Run() (int, error) {
	files, err := ioutil.ReadDir(aiffDir)
	if err != nil {
		return 0, errors.Wrap(err, "failed to read aiff recording dir")
	}

	var conversionCount int

	for _, f := range files {
		if !isAiff(f) {
			continue
		}

		aiff := fmt.Sprintf("%s%s", aiffDir, f.Name())
		flac := fmt.Sprintf("%s%s", flacDir, f.Name())
		flac = strings.Replace(flac, ".aiff", ".flac", 1)

		if flacExists(flac) {
			continue
		}

		if err := exec.Command("ffmpeg", "-i", aiff, "-c:a", "flac", flac).Run(); err != nil {
			return conversionCount, errors.Wrap(err, "error occured converting aiff files to flac")
		}

		fmt.Println("Converted: " + aiff + " to Flac:" + flac)
		conversionCount++
	}

	return conversionCount, nil
}

func isAiff(f os.FileInfo) bool {
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
