package postgresql

type TokenRepository struct {
	store *Store
}

func (r *TokenRepository) Save(userID int, refreshToken string) error {
	_, err := r.store.db.Exec(
		`INSERT INTO refresh_tokens (user_id, token) 
		VALUES ($1, $2) 
		ON CONFLICT (user_id) 
		DO UPDATE SET token = $2`,
		userID, refreshToken,
	)
	return err
}

func (r *TokenRepository) GetByUserID(userID int) (string, error) {
	var refreshToken string
	err := r.store.db.Get(&refreshToken, "SELECT token FROM refresh_tokens WHERE user_id=$1", userID)
	return refreshToken, err
}

func (r *TokenRepository) Delete(userID int) error {
	_, err := r.store.db.Exec("DELETE FROM refresh_tokens WHERE user_id = $1", userID)
	return err
}
