package models

type Comment struct {
	ID         int       `db:"id" json:"id"`
	PostID     int       `db:"post_id" json:"postID"`
	AuthorID   int       `db:"author_id" json:"authorID"`
	ParentID   *int      `db:"parent_id" json:"parentID"`
	AuthorName string    `db:"author_name" json:"authorName"`
	Text       string    `db:"text" json:"text"`
	CreatedAt  string    `db:"created_at" json:"createdAt"`
	Children   []Comment `json:"children"`
}
