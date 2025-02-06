package postgresql

import (
	"github.com/Crabocod/gpt_network/api-service/internal/models"
)

type CommentRepository struct {
	store *Store
}

func (r *CommentRepository) GetList(postID, offset, recordsPerPage int) ([]models.Comment, error) {
	var comments []models.Comment
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
	err := r.store.db.Select(&comments, query, postID, recordsPerPage, offset)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepository) Save(c models.Comment) error {
	var err error
	if c.ID > 0 {
		query := `UPDATE comments SET text = $1 WHERE id = $2`
		_, err = r.store.db.Exec(query, c.Text, c.ID)
	} else {
		if c.ParentID != nil {
			query := `INSERT INTO comments (post_id, author_id, parent_id, text) VALUES ($1, $2, $3, $4) RETURNING id`
			err = r.store.db.QueryRow(query, c.PostID, c.AuthorID, c.ParentID, c.Text).Scan(&c.ID)
		} else {
			query := `INSERT INTO comments (post_id, author_id, text) VALUES ($1, $2, $3) RETURNING id`
			err = r.store.db.QueryRow(query, c.PostID, c.AuthorID, c.Text).Scan(&c.ID)
		}
	}
	return err
}

func (r *CommentRepository) Delete(c models.Comment) error {
	_, err := r.store.db.Exec("DELETE FROM comments WHERE id = $1", c.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *CommentRepository) GetCount(postID int) (int, error) {
	var count int
	err := r.store.db.QueryRow("SELECT COUNT(*) FROM comments WHERE post_id = $1", postID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
