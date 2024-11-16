package models

import (
	"web.app/internal/db"
)

type Comment struct {
	ID        int    `db:"id" json:"id"`
	AuthorID  string `db:"author_id" json:"authorID"`
	PostID    string `db:"post_id" json:"postID"`
	ParentID  string `db:"parent_id" json:"parentID"`
	Text      string `db:"text" json:"text"`
	CreatedAt string `db:"created_at" json:"createdAt"`
}

func CreateComment(post_id int, author_id int, parent_id int, text string) error {
	_, err := db.DB.Exec("INSERT INTO comments (post_id, author_id, parent_id, text) VALUES ($1, $2, $3, $4)", post_id, author_id, author_id, text)
	if err != nil {
		return err
	}
	return nil
}
