package service

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
)

const (
	aiffExt = "aiff"
)

// AiffEncoder ...
type AiffEncoder struct {
	data []int32
}

// NewAiffEncoder ...
func NewAiffEncoder() *AiffEncoder {
	return &AiffEncoder{}
}

// Append ...
func (ae *AiffEncoder) Append(data []int32) {
	ae.data = append(ae.data, data...)
}

// SampleCount ...
func (ae *AiffEncoder) SampleCount() int {
	return len(ae.data)
}

// Encode ...
func (ae *AiffEncoder) Encode() (*bytes.Buffer, error) {
	var b bytes.Buffer
	f := bufio.NewWriter(&b)
	c := ae.SampleCount()

	if c == 0 {
		return nil, errors.New("no data to encode")
	}

	// form chunk
	_, err := f.WriteString("FORM")
	if err != nil {
		return nil, err
	}

	binary.Write(f, binary.BigEndian, ae.formSize()) //total bytes
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
	binary.Write(f, binary.BigEndian, int32(c))                        //number of samples
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

	binary.Write(f, binary.BigEndian, ae.soundSize()) //size
	binary.Write(f, binary.BigEndian, int32(0))       //offset
	binary.Write(f, binary.BigEndian, int32(0))       //block
	binary.Write(f, binary.BigEndian, ae.data)

	return &b, nil
}

func (ae *AiffEncoder) formSize() int32 {
	return int32(4 + 8 + 18 + 8 + 8 + 4*ae.SampleCount())
}

func (ae *AiffEncoder) soundSize() int32 {
	return int32(4*ae.SampleCount() + 8)
}
