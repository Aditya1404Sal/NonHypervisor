package main

import (
	"fmt"
	"log"
	"os"

	"Nonhypervisor/internal/builder"
	"Nonhypervisor/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: mockdocker <path to config file>")
	}

	configFilePath := os.Args[1]

	// Parse the configuration file
	cfg, err := config.ParseConfig(configFilePath)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	// Build the image based on the configuration
	err = builder.BuildImage(cfg)
	if err != nil {
		log.Fatalf("Error building image: %v", err)
	}

	fmt.Println("Image built successfully!")
}
