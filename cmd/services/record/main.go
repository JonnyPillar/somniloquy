package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/jonnypillar/somniloquy/internal/service/record"
	"google.golang.org/grpc"
)

const (
	fileSaverKey = "file"
	s3SaverKey   = "s3"
)

func main() {
	fmt.Println("Starting Up Somoiloquy Record Service")

	config, err := config.NewServiceConfig()
	if err != nil {
		log.Fatal("error occured creating config", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	grpcServer := grpc.NewServer()
	s, err := savers(config)
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	record.NewRecordingService(config, grpcServer, s...)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Something went wrong", err)
	}

	fmt.Println("Completed Record Service")
}

func savers(config *config.ServiceConfig) ([]record.Saver, error) {
	savers := []record.Saver{}

	for _, i := range config.UploadDestinations {
		if i == fileSaverKey {
			savers = append(savers, record.NewFile(config))
		} else if i == s3SaverKey {
			s3, err := record.NewS3Bucket(config)
			if err != nil {
				return nil, err
			}

			savers = append(savers, s3)
		}
	}

	return savers, nil
}
