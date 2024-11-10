package handlers

import (
	"encoding/json"
	"net/http"

	"web.app/internal/models"
	"web.app/internal/services"

	"github.com/golang-jwt/jwt/v5"
)

var creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&creds)

	err := models.RegisterUser(creds.Username, services.HashPassword(creds.Password))
	if err != nil {
		http.Error(w, `{"error": "Failed to register user"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User registered successfully"}`))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user_id")

	user, err := models.GetUserByID(user_id)
	if err != nil {
		http.Error(w, `{"error": "User not found"}`, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
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
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		RefreshToken string `json:"refresh_token"`
	}
	json.NewDecoder(r.Body).Decode(&requestData)

	claims := &services.JWTClaims{}
	token, err := jwt.ParseWithClaims(requestData.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return services.RefreshSecret, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, `{"error": "Invalid or expired refresh token"}`, http.StatusUnauthorized)
		return
	}

	// Проверка refresh_token в базе данных
	savedRefreshToken, err := models.GetRefreshToken(claims.UserID)
	if err != nil || savedRefreshToken != requestData.RefreshToken {
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
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
