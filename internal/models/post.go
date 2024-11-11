package models

import (
	"web.app/internal/db"
)

type Post struct {
	ID        int    `db:"id" json:"id"`
	AuthorID  string `db:"author_id" json:"authorID"`
	Text      string `db:"text" json:"text"`
	CreatedAt string `db:"created_at" json:"createdAt"`
}

func GetPosts(offset, recordsPerPage int) ([]Post, error) {
	var posts []Post
	query := `SELECT id, author_id, text, created_at FROM posts ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	err := db.DB.Select(&posts, query, recordsPerPage, offset)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func CountPosts() (int, error) {
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM posts").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CreatePost(author_id int, text string) error {
	_, err := db.DB.Exec("INSERT INTO posts (author_id, text) VALUES ($1, $2)", author_id, text)
	if err != nil {
		return err
	}
	return nil
}

func UpdatePost(author_id int, post_id, text string) error {
	query := `UPDATE posts SET text = $1 WHERE id = $2 AND author_id = $3`
	_, err := db.DB.Exec(query, text, post_id, author_id)
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(author_id int, post_id string) error {
	_, err := db.DB.Exec("DELETE FROM posts WHERE author_id = $1 AND id = $2", author_id, post_id)
	if err != nil {
		return err
	}
	return nil
}
