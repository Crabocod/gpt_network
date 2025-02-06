package models

type RefreshToken struct {
	ID     int    `db:"id" json:"id"`
	UserID int    `db:"user_id"  json:"userID"`
	Token  string `db:"token"  json:"refreshToken"`
}
