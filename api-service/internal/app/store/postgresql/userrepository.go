package postgresql

import (
	"github.com/Crabocod/gpt_network/api-service/internal/db"
	"github.com/Crabocod/gpt_network/api-service/internal/logger"
	"github.com/Crabocod/gpt_network/api-service/internal/models"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Save(user models.User) error {
	_, err := r.store.db.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", user.Username, user.PasswordHash)
	if err != nil {
		logger.Logrus.Fatalf("Error registering user: %v", err)
		return err
	}
	return nil
}

func (r *UserRepository) Get(username, passwordHash string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT * FROM users WHERE username = $1 AND password_hash = $2"
	err := r.store.db.Get(user, query, username, passwordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByID(id int) (*models.User, error) {
	var user models.User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
