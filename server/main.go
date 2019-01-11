package main

import (
	"fmt"
	"log"
	"net"

	"github.com/jonnypillar/somniloquy/api"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	s := api.Server{}

	grpcServer := grpc.NewServer()
	api.RegisterPingServer(grpcServer, &s)

	api.RegisterUploadServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Something went wrong", err)
	}
}
