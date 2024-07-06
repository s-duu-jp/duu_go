package main

import (
	"fmt"
	"log"

	"api/config/env"
)

func main() {
	// Load configuration
	cfg, err := env.GetConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}
	fmt.Println("env:", cfg["ENV"])
}
