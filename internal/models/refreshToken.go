package models

import (
	"web.app/internal/db"
)

func SaveRefreshToken(userID int, refreshToken string) error {
	_, err := db.DB.Exec(
		`INSERT INTO refresh_tokens (user_id, token) 
		VALUES ($1, $2) 
		ON CONFLICT (user_id) 
		DO UPDATE SET token = $2`,
		userID, refreshToken,
	)
	return err
}

func GetRefreshToken(userID int) (string, error) {
	var refreshToken string
	err := db.DB.Get(&refreshToken, "SELECT token FROM refresh_tokens WHERE user_id=$1", userID)
	return refreshToken, err
}

func DeleteRefreshToken(userID int) error {
	_, err := db.DB.Exec("DELETE FROM refresh_tokens WHERE user_id = $1", userID)
	return err
}
