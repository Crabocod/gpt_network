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

	generatePost()
	generateComment()
}

func generateComment() {
	var err error
	var comment handlers.Comment

	comment.AuthorName = services.RandomChoice(models)

	post, err := handlers.GetPost(comment.AuthorName)
	if err != nil {
		log.Fatalf("Ошибка при получении поста для автора '%s': %v", comment.AuthorName, err)
	}
	if post == nil {
		log.Fatalf("Пост для автора '%s' не найден", comment.AuthorName)
	}

	comment.Text, err = handlers.GenerateText(post.Text, comment.AuthorName)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	comment.PostID = post.ID
	err = comment.Save()
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	fmt.Println("Ответ от Python: ", post)
}

func generatePost() {
	var err error
	var post handlers.Post
	question := services.RandomChoice(questions)
	post.AuthorName = services.RandomChoice(models)

	post.Text, err = handlers.GenerateText(question, post.AuthorName)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	err = post.Save()
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	fmt.Println("Ответ от Python: ", post.Text)
}
