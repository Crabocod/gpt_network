package store

import "github.com/Crabocod/gpt_network/api-service/internal/models"

type UserRepository interface {
	Save(user models.User) error
	Get(username, passwordHash string) (*models.User, error)
	GetByID(id int) (*models.User, error)
	GetUserByName(name string) (*models.User, error)
}

type TokenRepository interface {
	Save(userID int, refreshToken string) error
	GetByUserID(userID int) (string, error)
	Delete(userID int) error
}

type PostRepository interface {
	GetList(offset, recordsPerPage int) ([]models.Post, error)
	GetCount() (int, error)
	Save(p models.Post) error
	Delete(p models.Post) error
	GetLatestFilteredPost(excludedAuthorName string) (*models.Post, error)
}

type CommentRepository interface {
	GetList(postID, offset, recordsPerPage int) ([]models.Comment, error)
	Save(c models.Comment) error
	Delete(c models.Comment) error
	GetCount(postID int) (int, error)
}
