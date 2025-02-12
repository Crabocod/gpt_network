package main

import (
	"github.com/Crabocod/gpt_network/api-service/cmd/factory"
	"log"
)

func main() {
	application, err := factory.InitializeApp()
	if err != nil {
		log.Fatal(err)
	}

	if err := application.Run(); err != nil {
		log.Fatal(err)
	}
}
