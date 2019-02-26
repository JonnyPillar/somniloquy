package client

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/gordonklaus/portaudio"
	"github.com/jonnypillar/somniloquy/config"
	"github.com/pkg/errors"
)

// Microphone ...
type Microphone struct {
	stream  *portaudio.Stream
	results []int32
}

// NewMicrophone ...
func NewMicrophone(config *config.ClientConfig) (*Microphone, error) {
	m := Microphone{}
	m.results = make([]int32, 64)

	portaudio.Initialize()
	audioStream, err := portaudio.OpenDefaultStream(1, 0, config.SampleRate, len(m.results), &m.results)
	if err != nil {
		return nil, errors.Wrap(err, "error occurred creating new microphone input")
	}

	m.stream = audioStream

	return &m, nil
}

// Start ...
func (m *Microphone) Start() {
	m.stream.Start()
}

// Read ...
func (m *Microphone) Read(p []byte) (int, error) {
	m.stream.Read()

	buf := &bytes.Buffer{}
	for _, v := range m.results {
		binary.Write(buf, binary.LittleEndian, v)
	}

	p = buf.Bytes()
	fmt.Println("buffer:", p[:10], "...", p[len(p)-10:])
	return len(p), nil
}

// ReadData ...
func (m *Microphone) ReadData() []int32 {
	m.stream.Read()
	return m.results
}

// Close ...
func (m *Microphone) Close() {
	m.stream.Close()
	portaudio.Terminate()
}
