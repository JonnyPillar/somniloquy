package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jonnypillar/somniloquy/config"
	"github.com/jonnypillar/somniloquy/internal/service"
	"github.com/jonnypillar/somniloquy/internal/service/filesystem"
)

func main() {
	fmt.Println("Starting Up Somoiloquy Transcription Service")

	config, err := config.NewServiceConfig()
	if err != nil {
		log.Fatal("error occured creating config", err)
	}

	r, err := filesystem.GetReader(config)
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	ctx := context.Background()
	gss, err := service.NewGoogleSpeechService(ctx)
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	ts := service.NewTranscriptionService(config, r, gss)
	results, err := ts.Start(ctx)
	if err != nil {
		log.Fatal("error occured transcribing", err)
	}

	fmt.Println(results)
	fmt.Println("Completed Transcription Service")
}
