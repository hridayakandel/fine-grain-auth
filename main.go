package main

import (
	"github.com/hridayakandel/fine-grain-auth/cmd"
	"log"
)

func main() {
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Error failing to start server: %s", err)
	}
	log.Println("fine-grain-auth service shutdown: Application stopped")
}
