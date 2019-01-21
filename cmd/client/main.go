package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/jonnypillar/somniloquy/internal/api"
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

	c, err := client.NewClient(config, conn)
	if err != nil {
		log.Fatal("Error occured creating client", err)
	}

	asc := api.NewAudioServiceClient(conn)

	stream, err := asc.Upload(context.Background())
	if err != nil {
		log.Fatal(err, "an error occured creating upload stream")
	}

	err = c.Stream(stream)
	if err != nil {
		log.Fatal(err, "an error occured creating upload stream")
	}
}
