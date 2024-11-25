package models

import (
	"web.app/internal/db"
)

type Post struct {
	ID         int     `db:"id" json:"id"`
	AuthorID   int     `db:"author_id" json:"authorID"`
	AuthorName string  `db:"author_name" json:"authorName"`
	Text       string  `db:"text" json:"text"`
	Photo      *string `db:"photo" json:"photo"`
	CreatedAt  string  `db:"created_at" json:"createdAt"`
}

func (p *Post) GetList(offset, recordsPerPage int) ([]Post, error) {
	var posts []Post
	query := `
		SELECT 
			p.*, 
			u.username AS author_name
		FROM 
			posts p
		JOIN 
			users u
		ON 
			p.author_id = u.id
		ORDER BY 
			p.created_at DESC 
		LIMIT $1 OFFSET $2`
	err := db.DB.Select(&posts, query, recordsPerPage, offset)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *Post) Count() (int, error) {
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM posts").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (p *Post) Save() error {
	if p.ID > 0 {
		query := `UPDATE posts SET text = $1 WHERE id = $2`
		_, err := db.DB.Exec(query, p.Text, p.ID)
		if err != nil {
			return err
		}
	} else {
		query := `INSERT INTO posts (author_id, text) VALUES ($1, $2) RETURNING id`
		err := db.DB.QueryRow(query, p.AuthorID, p.Text).Scan(&p.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Post) Delete() error {
	_, err := db.DB.Exec("DELETE FROM posts WHERE id = $1", p.ID)
	if err != nil {
		return err
	}
	return nil
}
