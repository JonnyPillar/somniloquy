package files

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

const (
	aiffExt = "aiff"
)

// Aiff ...
type Aiff struct {
	Filename string
	data     []int32
	samples  int
}

// NewAiff ...
func NewAiff() *Aiff {
	return &Aiff{
		Filename: fmt.Sprintf("%s.%s", time.Now().Format("2006-01-02 15:04:05"), aiffExt),
	}
}

// Append ...
func (r *Aiff) Append(content []int32) {
	r.data = append(r.data, content...)
	r.samples += len(content)

	fmt.Println("Stream Received", r.samples)
}

// Buffer ...
func (r *Aiff) Buffer() (*bytes.Buffer, error) {
	var b bytes.Buffer
	f := bufio.NewWriter(&b)

	// form chunk
	_, err := f.WriteString("FORM")
	if err != nil {
		return nil, err
	}

	totalBytes := 4 + 8 + 18 + 8 + 8 + 4*r.samples
	binary.Write(f, binary.BigEndian, int32(totalBytes)) //total bytes
	_, err = f.WriteString("AIFF")
	if err != nil {
		return nil, err
	}

	// common chunk
	_, err = f.WriteString("COMM")
	if err != nil {
		return nil, err
	}
	binary.Write(f, binary.BigEndian, int32(18))                       //size
	binary.Write(f, binary.BigEndian, int16(1))                        //channels
	binary.Write(f, binary.BigEndian, int32(r.samples))                //number of samples
	binary.Write(f, binary.BigEndian, int16(32))                       //bits per sample
	_, err = f.Write([]byte{0x40, 0x0e, 0xac, 0x44, 0, 0, 0, 0, 0, 0}) //80-bit sample rate 44100
	if err != nil {
		return nil, err
	}

	// sound chunk
	_, err = f.WriteString("SSND")
	if err != nil {
		return nil, err
	}

	binary.Write(f, binary.BigEndian, int32(4*r.samples+8)) //size
	binary.Write(f, binary.BigEndian, int32(0))             //offset
	binary.Write(f, binary.BigEndian, int32(0))             //block
	binary.Write(f, binary.BigEndian, r.data)

	return &b, nil
}
