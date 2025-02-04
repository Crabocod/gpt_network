package postgresql

import (
	"github.com/Crabocod/gpt_network/api-service/internal/db"
	"github.com/Crabocod/gpt_network/api-service/internal/models"
)

type PostRepository struct {
	store *Store
}

func (r *PostRepository) GetList(offset, recordsPerPage int) ([]models.Post, error) {
	var posts []models.Post
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

func (r *PostRepository) GetCount() (int, error) {
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM posts").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *PostRepository) Save(p models.Post) error {
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

func (r *PostRepository) Delete(p models.Post) error {
	_, err := db.DB.Exec("DELETE FROM posts WHERE id = $1", p.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) GetLatestFilteredPost(excludedAuthorName string) (*models.Post, error) {
	var post models.Post
	query := `
		SELECT 
			*
		FROM 
			posts
		WHERE 
			author_id != (
				SELECT id FROM users WHERE username = $1
			)
			AND NOT EXISTS (
				SELECT 1
				FROM comments c
				JOIN users cu ON c.author_id = cu.id
				WHERE 
					c.post_id = posts.id 
					AND c.parent_id IS NULL 
					AND cu.username = $1
			)
		ORDER BY 
			created_at DESC
		LIMIT 1`
	err := db.DB.Get(&post, query, excludedAuthorName)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
