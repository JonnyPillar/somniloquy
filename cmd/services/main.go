package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/jonnypillar/somniloquy/internal/service"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Up Somoiloquy Service")

	config, err := config.NewServiceConfig()
	if err != nil {
		log.Fatal("error occured creating config", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	grpcServer := grpc.NewServer()
	service.NewAudioService(config, grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Something went wrong", err)
	}
}
