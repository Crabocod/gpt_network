package main

import (
	"fmt"
	"log"

	"generate/internal/handlers"
)

func main() {
	text, err := handlers.Generate()
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	err = handlers.Save(text)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	fmt.Printf("Ответ от Python: %s\n", text)
}
