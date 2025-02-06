package models

type Post struct {
	ID         int     `db:"id" json:"id"`
	AuthorID   int     `db:"author_id" json:"authorID"`
	AuthorName string  `db:"author_name" json:"authorName"`
	Text       string  `db:"text" json:"text"`
	Photo      *string `db:"photo" json:"photo"`
	CreatedAt  string  `db:"created_at" json:"createdAt"`
}
