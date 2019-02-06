package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/jonnypillar/somniloquy/config"
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

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, os.Interrupt)

	defer close(cancelChan)
	defer func() {
		signal.Stop(cancelChan)
		cancel()
	}()

	go func() {
		select {
		case <-cancelChan:
			cancel()
		case <-ctx.Done():
		}
	}()

	c, err := client.NewClient(config, conn)
	if err != nil {
		log.Fatal("Error occured creating client", err)
	}

	asc := api.NewRecordServiceClient(conn)

	stream, err := asc.Upload(context.Background())
	if err != nil {
		log.Fatal(err, "an error occured creating upload stream")
	}

	err = c.Stream(ctx, stream)
	if err != nil {
		log.Fatal(err, "an error occured creating upload stream")
	}
}
