package main

import (
	"fmt"
	"log"

	"github.com/jonnypillar/somniloquy/internal/service/conversion"
)

func main() {
	fmt.Println("Starting Up Conversion Service")

	count, err := conversion.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Completed Conversion Service. Converted:", count)
}
