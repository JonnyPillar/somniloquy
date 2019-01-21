package service

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jonnypillar/somniloquy/internal/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const fileFormat = "./assets/recordings/%s.aiff"

// AudioService defines the Audio gRPC Service
type AudioService struct{}

// RegisterAudioService registers the Audio Service with the gRPC Server
func RegisterAudioService(grpcServer *grpc.Server) {
	as := &AudioService{}

	api.RegisterAudioServiceServer(grpcServer, as)
}

// Upload ...
func (s *AudioService) Upload(stream api.AudioService_UploadServer) error {
	f, err := os.Create(fmt.Sprintf(fileFormat, time.Now().Format("2006-01-02 15:04:05")))
	if err != nil {
		panic(err)
	}

	// form chunk
	_, err = f.WriteString("FORM")
	chk(err)
	chk(binary.Write(f, binary.BigEndian, int32(0))) //total bytes
	_, err = f.WriteString("AIFF")
	chk(err)

	// common chunk
	_, err = f.WriteString("COMM")
	chk(err)
	chk(binary.Write(f, binary.BigEndian, int32(18)))                  //size
	chk(binary.Write(f, binary.BigEndian, int16(1)))                   //channels
	chk(binary.Write(f, binary.BigEndian, int32(0)))                   //number of samples
	chk(binary.Write(f, binary.BigEndian, int16(32)))                  //bits per sample
	_, err = f.Write([]byte{0x40, 0x0e, 0xac, 0x44, 0, 0, 0, 0, 0, 0}) //80-bit sample rate 44100
	chk(err)

	// sound chunk
	_, err = f.WriteString("SSND")
	chk(err)
	chk(binary.Write(f, binary.BigEndian, int32(0))) //size
	chk(binary.Write(f, binary.BigEndian, int32(0))) //offset
	chk(binary.Write(f, binary.BigEndian, int32(0))) //block
	nSamples := 0

	defer func() {
		// fill in missing sizes
		totalBytes := 4 + 8 + 18 + 8 + 8 + 4*nSamples
		_, err = f.Seek(4, 0)
		chk(err)
		chk(binary.Write(f, binary.BigEndian, int32(totalBytes)))
		_, err = f.Seek(22, 0)
		chk(err)
		chk(binary.Write(f, binary.BigEndian, int32(nSamples)))
		_, err = f.Seek(42, 0)
		chk(err)
		chk(binary.Write(f, binary.BigEndian, int32(4*nSamples+8)))
		chk(f.Close())
	}()

	// check(err)

	for {
		c, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}

			err = errors.Wrapf(err,
				"failed unexpectadely while reading chunks from stream")
			return err
		}

		fmt.Println("Stream Received", nSamples)
		binary.Write(f, binary.BigEndian, c.Content)
		nSamples += len(c.Content)
	}

	err = stream.SendAndClose(&api.UploadStatus{
		Message: "Upload received with success",
		Code:    api.UploadStatusCode_Ok,
	})
	if err != nil {
		err = errors.Wrapf(err,
			"failed to send status code")
		return err
	}

	log.Println("Messaged Reseived", nSamples)

	return err
}
func chk(err error) {
	if err != nil {
		panic(err)
	}
}
