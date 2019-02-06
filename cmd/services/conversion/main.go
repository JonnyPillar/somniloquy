package main

import (
	"fmt"
	"log"

	"github.com/jonnypillar/somniloquy/config"
	"github.com/jonnypillar/somniloquy/internal/service"
	"github.com/jonnypillar/somniloquy/internal/service/filesystem"
)

const (
	fileReaderKey = "file"
	s3ReaderKey   = "s3"
)

func main() {
	fmt.Println("Starting Up Conversion Service")

	config, err := config.NewServiceConfig()
	if err != nil {
		log.Fatal("error occured creating config", err)
	}

	r, err := filesystem.GetReader(config)
	if err != nil {
		log.Fatal("Something went wrong", err)
	}

	converter := service.FFMPEGConverter{}

	afc := service.NewAIFFConverter(config, r, converter)

	count, err := afc.ToFlac()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Completed Conversion Service. Converted:", count)
}
