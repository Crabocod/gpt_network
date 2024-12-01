package models

import (
	"log"

	"web.app/internal/db"
)

type User struct {
	ID           int    `db:"id" json:"id"`
	Username     string `db:"username" json:"username"`
	PasswordHash string `db:"password_hash" json:"password"`
	CreatedAt    string `db:"created_at" json:"date"`
}

func (u *User) Register() error {
	_, err := db.DB.Exec("INSERT INTO users (username, password_hash) VALUES ($1, $2)", u.Username, u.PasswordHash)
	if err != nil {
		log.Printf("Error registering user: %v", err)
		return err
	}
	return nil
}

func (u *User) Login() error {
	err := db.DB.Get(u, "SELECT * FROM users WHERE username=$1 AND password_hash=$2", u.Username, u.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(id int) (*User, error) {
	var user User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByName(name string) (*User, error) {
	var user User
	err := db.DB.Get(&user, "SELECT * FROM users WHERE username=$1", name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
