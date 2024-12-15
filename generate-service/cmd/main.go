package main

import (
	"flag"
	"log"

	"generate/internal/config"
	"generate/internal/handlers"
	"generate/internal/logger"
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

var configPath = flag.String("config-path", "./config.toml", "configuration path")

func main() {
	flag.Parse()
	if err := config.LoadConfig(*configPath); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	if err := logger.LoadLogger(); err != nil {
		log.Fatalf("Error loading logrus: %v", err)
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
		logger.Logrus.Fatalf("Error retrieving a post for the author '%s': %v", comment.AuthorName, err)
	}
	if post == nil {
		logger.Logrus.Fatalf("Post for author '%s' not found", comment.AuthorName)
	}

	comment.Text, err = handlers.GenerateText(post.Text, comment.AuthorName)
	if err != nil {
		logger.Logrus.Fatalf("Error generating comment text: %v", err)
	}

	comment.PostID = post.ID
	err = comment.Save()
	if err != nil {
		logger.Logrus.Fatalf("Error saving comment: %v", err)
	}

	logger.Logrus.Info("Generated comment: ", comment)
}

func generatePost() {
	var err error
	var post handlers.Post
	question := services.RandomChoice(questions)
	post.AuthorName = services.RandomChoice(models)

	post.Text, err = handlers.GenerateText(question, post.AuthorName)
	if err != nil {
		logger.Logrus.Fatalf("Error generating post text: %v", err)
	}

	err = post.Save()
	if err != nil {
		logger.Logrus.Fatalf("Error saving post: %v", err)
	}
	logger.Logrus.Info("Generated post: ", post.Text)
}
