package handlers

import (
	"net/http"
)

type Comment struct {
	ID       string `json:"id"`
	AuthorID string `json:"authorId"`
	Text     string `json:"text"`
	Date     string `json:"date"`
}

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
}
