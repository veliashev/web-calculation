package main

import (
	"log"

	"github.com/veliashev/web-calculation/config"
	"github.com/veliashev/web-calculation/internal/application"
)

func main() {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error getting config: %s", err)
	}

	app := application.New(config.Port)
	err = app.Run()
	if err != nil {
		log.Fatalf("Error starting application: %s", err)
	}
}
