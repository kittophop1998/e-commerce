package main

import (
	"log"

	"github.com/kittiphop/e-commerce/internal/bootstrap"
)

func main() {
	// Initialize application
	app, err := bootstrap.InitializeApp()
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	error := app.Listen(":3001")
	if error != nil {
		log.Fatalf("Error starting server: %v", error)
	}
}
