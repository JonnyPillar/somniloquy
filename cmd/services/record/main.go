package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jonnypillar/somniloquy/config"
	"github.com/jonnypillar/somniloquy/internal/api"
	"github.com/jonnypillar/somniloquy/internal/service"
	"github.com/jonnypillar/somniloquy/internal/service/filesystem"
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
	s, err := filesystem.GetSaver(config)
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	rs := service.NewRecordingService(config, s)

	api.RegisterRecordServiceServer(grpcServer, rs)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Something went wrong", err)
	}

	fmt.Println("Completed Record Service")
}
