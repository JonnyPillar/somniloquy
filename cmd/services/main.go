package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jonnypillar/somniloquy/internal/service"
	"google.golang.org/grpc"
)

const port = 7777

func main() {
	fmt.Println("Starting Up Somoiloquy Service")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("Something went wrong", err)
	}
	grpcServer := grpc.NewServer()

	service.RegisterAudioService(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Something went wrong", err)
	}
}
