package models

type User struct {
	ID           int    `db:"id" json:"id"`
	Username     string `db:"username" json:"username"`
	Password     string `json:"password"`
	PasswordHash string `db:"password_hash"`
	CreatedAt    string `db:"created_at" json:"date"`
}
