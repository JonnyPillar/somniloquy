package fake

import "bytes"

// Saver ...
type Saver struct {
	Error error
}

// Save ...
func (s Saver) Save(filename string, by *bytes.Buffer) error {
	return s.Error
}
