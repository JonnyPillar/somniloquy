package main

import (
	"fmt"
	"log"

	"github.com/jonnypillar/somniloquy/configs"
	"github.com/jonnypillar/somniloquy/internal/service/transcription"
)

func main() {
	fmt.Println("Starting Up Somoiloquy Transcription Service")

	config, err := config.NewServiceConfig()
	if err != nil {
		log.Fatal("error occured creating config", err)
	}

	ts := transcription.NewTranscriptionService(config)
	err = ts.Start()
	if err != nil {
		log.Fatal("error occured transcribing", err)
	}

	fmt.Println("Completed Transcription Service")
}
