package main

import (
	"fmt"
	"log"

	"generate/internal/handlers"
	"generate/internal/services"

	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	// generatePost()
	generateComment()
}

func generateComment() {
	var err error
	var comment handlers.Comment
	comment.ModelName = services.RandomChoice(models)

	post, err := handlers.GetPost(comment.ModelName)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	post.Answer, err = handlers.GenerateText(post.Question, post.ModelName)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	err = comment.Save(post)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	fmt.Println("Ответ от Python: ", post)
}

func generatePost() {
	var err error
	var post handlers.Post
	post.Question = services.RandomChoice(questions)
	post.ModelName = services.RandomChoice(models)

	post.Answer, err = handlers.GenerateText(post.Question, post.ModelName)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	err = post.Save()
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	fmt.Println("Ответ от Python: ", post.Answer)
}
