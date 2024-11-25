package models

import (
	"web.app/internal/db"
)

type RefreshToken struct {
	ID     int    `db:"id" json:"id"`
	UserID int    `db:"user_id"  json:"userID"`
	Token  string `db:"token"  json:"refreshToken"`
}

func (r RefreshToken) Save() error {
	_, err := db.DB.Exec(
		`INSERT INTO refresh_tokens (user_id, token) 
		VALUES ($1, $2) 
		ON CONFLICT (user_id) 
		DO UPDATE SET token = $2`,
		r.UserID, r.Token,
	)
	return err
}

func (r RefreshToken) GetByUserID(userID int) (string, error) {
	var refreshToken string
	err := db.DB.Get(&refreshToken, "SELECT token FROM refresh_tokens WHERE user_id=$1", userID)
	return refreshToken, err
}

func (r RefreshToken) Delete() error {
	_, err := db.DB.Exec("DELETE FROM refresh_tokens WHERE user_id = $1", r.UserID)
	return err
}
