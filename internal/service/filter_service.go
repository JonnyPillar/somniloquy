package service

import (
	"github.com/jonnypillar/somniloquy/config"
	"github.com/pkg/errors"
)

// ReaderDeleter ...
type ReaderDeleter interface {
	Reader
	Delete() error
}

// FilterService ...
type FilterService struct {
	config *config.ServiceConfig
	r      Reader
}

// NewFilterService ...
func NewFilterService(c *config.ServiceConfig, r ReaderDeleter) *FilterService {
	return &FilterService{
		config: c,
		r:      r,
	}
}

// Filter ...
func (fs FilterService) Filter() error {
	files, err := fs.r.Read()
	if err != nil {
		return errors.Wrap(err, "failed to read aiff recording dir")
	}

	// activeFiles := []os.FileInfo{}

	for _, f := range files {
		if !isAiff(f) {
			continue
		}

		// aiff := fmt.Sprintf("%s%s", fs.config.AIFFRecordingFilePath, f.Name())

		// data, err := ioutil.ReadFile(aiff)
		// if err != nil {
		// 	return errors.Wrap(err, "failed to read flac recording")
		// }

		// engine, err := sed.New(ioutil.ReadFile(aiff))

		// output, err := engine.RunString(inString)

		// activeFiles = append(activeFiles, f)
	}

	return nil
}

// echo 2019-03-11 12:59:21.aiff | sed -n 's#^\([0-9]+\)-max-\([0-9.]+\)-mid-\([0-9.]+\)\.'"aiff"'$#\1#p'
