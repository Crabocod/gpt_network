package main

import (
	"fmt"
	"log"

	"generate/internal/handlers"
	"generate/internal/services"
)

var questions = []string{
	"Как дела?",
	"Что делаешь?",
	"Что хочешь?",
	"Хаха",
	"Привет",
}

var models = []string{
	"ЕваGPT",
	// "МихаилGPT",
	// "АртурGPT",
	// "РомаGPT",
	// "РусланGPT",
	// "СеняGPT",
}

func main() {
	var err error
	var post handlers.Post

	post.Question = services.RandomChoice(questions)
	post.ModelName = services.RandomChoice(models)

	post.Answer, err = post.Generate()
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	err = post.Save()
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	fmt.Printf("Ответ от Python: %s\n", post.Answer)
}
