package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"web.app/internal/models"
	"web.app/internal/services"

	"github.com/golang-jwt/jwt/v5"
)

type Creds struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	RefreshToken string `json:"refresh_token"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var creds Creds
	json.NewDecoder(r.Body).Decode(&creds)

	err := models.RegisterUser(creds.Username, services.HashPassword(creds.Password))
	if err != nil {
		http.Error(w, `{"error": "Failed to register user"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User registered successfully"}`))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var creds Creds
	json.NewDecoder(r.Body).Decode(&creds)

	user, err := models.GetUserByID(creds.ID)
	if err != nil {
		http.Error(w, `{"error": "User not found"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"id":       strconv.Itoa(user.ID),
		"username": user.Username,
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)

	err := models.DeleteRefreshToken(userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to log out user"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "User logged out successfully"}`))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Creds
	json.NewDecoder(r.Body).Decode(&creds)

	user, err := models.GetUserByUsernameAndPassword(creds.Username, services.HashPassword(creds.Password))
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

	refreshToken, err := services.GenerateRefreshToken(user.ID)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate refresh token"}`, http.StatusInternalServerError)
		return
	}

	// Сохранение refresh_token в базе данных
	err = models.SaveRefreshToken(user.ID, refreshToken)
	if err != nil {
		http.Error(w, `{"error": "Failed to save refresh token"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	var creds Creds
	json.NewDecoder(r.Body).Decode(&creds)

	claims := &services.JWTClaims{}
	token, err := jwt.ParseWithClaims(creds.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return services.RefreshSecret, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, `{"error": "Invalid or expired refresh token"}`, http.StatusUnauthorized)
		return
	}

	// Проверка refresh_token в базе данных
	savedRefreshToken, err := models.GetRefreshToken(claims.UserID)
	if err != nil || savedRefreshToken != creds.RefreshToken {
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
	refreshToken, err := services.GenerateRefreshToken(claims.UserID)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate refresh token"}`, http.StatusInternalServerError)
		return
	}

	// Сохранение refresh_token в базе данных
	err = models.SaveRefreshToken(claims.UserID, refreshToken)
	if err != nil {
		http.Error(w, `{"error": "Failed to save refresh token"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
