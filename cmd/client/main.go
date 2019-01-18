package main

import (
	"fmt"
	"log"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/jonnypillar/somniloquy/internal/client"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Client")

	config, err := config.NewClientConfig()
	if err != nil {
		log.Fatal("error occured creating config", err)
	}

	conn, err := grpc.Dial(config.ServiceURL(), grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect", err)
	}

	c := client.NewClient(config, conn)

	c.Send()
}
