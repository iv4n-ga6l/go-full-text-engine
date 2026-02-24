package main

import (
	"fmt"
	"log"

	"github.com/iv4n-ga6l/go-full-text-engine/index"
)

func main() {
	// Initialize the inverted and TF-IDF indexes
	if err := index.Init(); err != nil {
		log.Fatalf("Failed to initialize indexes: %v", err)
	}

	fmt.Println("Indexes initialized successfully. Application is ready.")

	// TODO: Add the rest of the application logic here
}