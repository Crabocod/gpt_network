package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"web.app/internal/models"
	"web.app/internal/services"

	"github.com/golang-jwt/jwt/v5"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}
	user.PasswordHash = services.HashPassword(user.PasswordHash)

	err := user.Register()
	if err != nil {
		http.Error(w, `{"error": "Failed to register user"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User registered successfully"}`))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	User, err := user.GetByID(user.ID)
	if err != nil {
		http.Error(w, `{"error": "User not found"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"id":       strconv.Itoa(User.ID),
		"username": User.Username,
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	var refreshToken models.RefreshToken
	refreshToken.UserID = r.Context().Value("user_id").(int)

	err := refreshToken.Delete()
	if err != nil {
		http.Error(w, `{"error": "Failed to log out user"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "User logged out successfully"}`))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var refreshToken models.RefreshToken
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}
	user.PasswordHash = services.HashPassword(user.PasswordHash)

	err := user.Login()
	if err != nil {
		http.Error(w, `{"error": "Invalid client credentials"}`, http.StatusUnauthorized)
		return
	}

	// Генерация access_token и refresh_token
	accessToken, err := services.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate access token"}`, http.StatusInternalServerError)
		return
	}

	refreshToken.UserID = user.ID
	refreshToken.Token, err = services.GenerateRefreshToken(user.ID)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate refresh token"}`, http.StatusInternalServerError)
		return
	}

	// Сохранение refresh_token в базе данных
	err = refreshToken.Save()
	if err != nil {
		http.Error(w, `{"error": "Failed to save refresh token"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken.Token,
	})
}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	var refreshToken models.RefreshToken
	json.NewDecoder(r.Body).Decode(&refreshToken)

	claims := &services.JWTClaims{}
	token, err := jwt.ParseWithClaims(refreshToken.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return services.RefreshSecret, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, `{"error": "Invalid or expired refresh token"}`, http.StatusUnauthorized)
		return
	}

	// Проверка refresh_token в базе данных
	savedRefreshToken, err := refreshToken.GetByUserID(claims.UserID)
	if err != nil || savedRefreshToken != refreshToken.Token {
		http.Error(w, `{"error": "Refresh token not recognized"}`, http.StatusUnauthorized)
		return
	}

	// Генерация нового access_token
	accessToken, err := services.GenerateJWT(claims.UserID)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate access token"}`, http.StatusInternalServerError)
		return
	}

	// Генерация нового refresh_token
	refreshToken.UserID = claims.UserID
	refreshToken.Token, err = services.GenerateRefreshToken(claims.UserID)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate refresh token"}`, http.StatusInternalServerError)
		return
	}

	// Сохранение refresh_token в базе данных
	err = refreshToken.Save()
	if err != nil {
		http.Error(w, `{"error": "Failed to save refresh token"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken.Token,
	})
}
