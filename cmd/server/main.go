package main

import (
	"log"

	"github.com/afiffaizun/mining-asset-management/internal/app"
)

func main() {
	application, err := app.New()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
