package main

import (
	"log"

	"github.com/kittiphop/e-commerce/internal/bootstrap"
)

func main() {
	/* Initialize application by gin */
	app, err := bootstrap.InitializeApp()
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	error := app.Run(":3100")
	if error != nil {
		log.Fatalf("Error starting server: %v", error)
	}
}
