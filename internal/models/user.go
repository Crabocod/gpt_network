package models

import (
	"log"

	"web.app/internal/db"
)

type User struct {
	ID           int    `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
	CreatedAt    string `db:"created_at"`
}

func RegisterUser(username, passwordHash string) error {
	_, err := db.DB.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", username, passwordHash)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		return err
	}
	return nil
}

func GetUserByUsernameAndPassword(username, passwordHash string) (*User, error) {
	var user User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE username=$1 AND password_hash=$2", username, passwordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByID(id interface{}) (*User, error) {
	var user User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
