package main

import (
	"log"

	"github.com/sahilrana7582/user-service/configs"
)

func main() {
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	log.Println("Configuration loaded successfully: %+v", cfg)
}
