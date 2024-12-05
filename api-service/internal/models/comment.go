package models

import (
	"web.app/internal/db"
)

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

func (c *Comment) GetList(postID, offset, recordsPerPage int) ([]Comment, error) {
	var comments []Comment
	query := `
		SELECT 
			c.*, 
			u.username AS author_name
		FROM 
			comments c
		JOIN 
			users u
		ON 
			c.author_id = u.id
		WHERE
			c.post_id = $1
		ORDER BY 
			c.created_at DESC 
		LIMIT $2 OFFSET $3`
	err := db.DB.Select(&comments, query, postID, recordsPerPage, offset)
	if err != nil {
		return nil, err
	}
	return buildTree(comments), nil
}

func (c *Comment) Save() error {
	var err error
	if c.ID > 0 {
		query := `UPDATE comments SET text = $1 WHERE id = $2`
		_, err = db.DB.Exec(query, c.Text, c.ID)
	} else {
		if c.ParentID != nil {
			query := `INSERT INTO comments (post_id, author_id, parent_id, text) VALUES ($1, $2, $3, $4) RETURNING id`
			err = db.DB.QueryRow(query, c.PostID, c.AuthorID, c.ParentID, c.Text).Scan(&c.ID)
		} else {
			query := `INSERT INTO comments (post_id, author_id, text) VALUES ($1, $2, $3) RETURNING id`
			err = db.DB.QueryRow(query, c.PostID, c.AuthorID, c.Text).Scan(&c.ID)
		}
	}
	return err
}

func (c *Comment) Delete() error {
	_, err := db.DB.Exec("DELETE FROM comments WHERE id = $1", c.ID)
	if err != nil {
		return err
	}
	return nil
}

func (c *Comment) Count(postID int) (int, error) {
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM comments WHERE post_id = $1", postID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func buildTree(comments []Comment) []Comment {
	commentMap := make(map[int]*Comment, len(comments))
	var roots []Comment

	for i := range comments {
		comment := &comments[i]
		comment.Children = make([]Comment, 0)
		commentMap[comment.ID] = comment
	}

	for i := range comments {
		comment := &comments[i]
		if comment.ParentID == nil {
			roots = append(roots, *comment)
		} else if parent, exists := commentMap[*comment.ParentID]; exists {
			parent.Children = append(parent.Children, *comment)
		}
	}

	return roots
}
