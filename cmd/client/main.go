package main

import (
	"fmt"
	"log"

	"github.com/jonnypillar/somniloquy/internal/client"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting Client")

	conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Did not connect", err)
	}

	c := client.NewClient(conn)

	c.Send()
}
