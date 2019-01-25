package record

import (
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

const (
	aiffExt = "aiff"
	aiffDir = "./assets/recordings/aiff/"
)

// Recording ...
type Recording struct {
	data    []int32
	samples int
	file    string
}

// NewRecording ...
func NewRecording() *Recording {
	return &Recording{
		file: fmt.Sprintf("%s%s.%s", aiffDir, time.Now().Format("2006-01-02 15:04:05"), aiffExt),
	}
}

// Append ...
func (r *Recording) Append(content []int32) {
	r.data = append(r.data, content...)
	r.samples += len(content)

	fmt.Println("Stream Received", r.samples)
}

// Save ...
func (r *Recording) Save() error {
	f, err := os.Create(r.file)
	if err != nil {
		return err
	}

	// form chunk
	_, err = f.WriteString("FORM")
	if err != nil {
		return err
	}

	binary.Write(f, binary.BigEndian, int32(0)) //total bytes
	_, err = f.WriteString("AIFF")
	if err != nil {
		return err
	}

	// common chunk
	_, err = f.WriteString("COMM")
	if err != nil {
		return err
	}
	binary.Write(f, binary.BigEndian, int32(18))                       //size
	binary.Write(f, binary.BigEndian, int16(1))                        //channels
	binary.Write(f, binary.BigEndian, int32(0))                        //number of samples
	binary.Write(f, binary.BigEndian, int16(32))                       //bits per sample
	_, err = f.Write([]byte{0x40, 0x0e, 0xac, 0x44, 0, 0, 0, 0, 0, 0}) //80-bit sample rate 44100
	if err != nil {
		return err
	}

	// sound chunk
	_, err = f.WriteString("SSND")
	if err != nil {
		return err
	}

	binary.Write(f, binary.BigEndian, int32(0)) //size
	binary.Write(f, binary.BigEndian, int32(0)) //offset
	binary.Write(f, binary.BigEndian, int32(0)) //block

	defer func() {
		// fill in missing sizes
		totalBytes := 4 + 8 + 18 + 8 + 8 + 4*r.samples
		_, err = f.Seek(4, 0)
		binary.Write(f, binary.BigEndian, int32(totalBytes))
		_, err = f.Seek(22, 0)
		binary.Write(f, binary.BigEndian, int32(r.samples))
		_, err = f.Seek(42, 0)
		binary.Write(f, binary.BigEndian, int32(4*r.samples+8))
		f.Close()
	}()

	binary.Write(f, binary.BigEndian, r.data)

	fmt.Println("Saved recording to ", r.file)

	return nil
}
