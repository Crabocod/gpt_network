package rest

import (
	"encoding/json"
	"github.com/Crabocod/gpt_network/api-service/internal/app/service"
	"github.com/Crabocod/gpt_network/api-service/internal/models"
	"net/http"
	"strconv"
)

type UserControllerInterface interface {
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	LoginHandler(w http.ResponseWriter, r *http.Request)
	RefreshTokenHandler(w http.ResponseWriter, r *http.Request)
	LogoutHandler(w http.ResponseWriter, r *http.Request)
	GetUserHandler(w http.ResponseWriter, r *http.Request)
}

type UserController struct {
	service service.Service
}

func NewUserController(s service.Service) UserControllerInterface {
	return &UserController{
		service: s,
	}
}

func (c *UserController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	err := c.service.UserService.Save(user)
	if err != nil {
		http.Error(w, `{"error": "Failed to register user"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte(`{"message": "User registered successfully"}`))
	if err != nil {
		return
	}
}

func (c *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	user, err := c.service.UserService.Get(requestData.Username, requestData.Password)
	if err != nil {
		http.Error(w, `{"error": "Invalid client credentials"}`, http.StatusUnauthorized)
		return
	}

	accessToken, err := c.service.TokenService.GenerateAccess(user.ID)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate access token"}`, http.StatusInternalServerError)
		return
	}

	refreshToken, err := c.service.TokenService.GenerateRefresh(user.ID)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate refresh token"}`, http.StatusInternalServerError)
		return
	}

	err = c.service.TokenService.Save(user.ID, refreshToken)
	if err != nil {
		http.Error(w, `{"error": "Failed to save refresh token"}`, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
	if err != nil {
		return
	}
}

func (c *UserController) RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	var refreshToken models.RefreshToken
	err := json.NewDecoder(r.Body).Decode(&refreshToken)
	if err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	userID, err := c.service.UserService.GetIDByToken(refreshToken.Token)
	if err != nil {
		http.Error(w, `{"error": "Invalid or expired refresh token"}`, http.StatusUnauthorized)
		return
	}

	savedRefreshToken, err := c.service.TokenService.GetByUserID(userID)
	if err != nil || savedRefreshToken != refreshToken.Token {
		http.Error(w, `{"error": "Refresh token not recognized"}`, http.StatusUnauthorized)
		return
	}

	accessToken, err := c.service.TokenService.GenerateAccess(userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate access token"}`, http.StatusInternalServerError)
		return
	}

	refreshToken.Token, err = c.service.TokenService.GenerateRefresh(userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to generate refresh token"}`, http.StatusInternalServerError)
		return
	}

	err = c.service.TokenService.Save(userID, refreshToken.Token)
	if err != nil {
		http.Error(w, `{"error": "Failed to save refresh token"}`, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken.Token,
	})
	if err != nil {
		return
	}
}

func (c *UserController) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)

	err := c.service.TokenService.Delete(userID)
	if err != nil {
		http.Error(w, `{"error": "Failed to log out user"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(`{"message": "User logged out successfully"}`))
	if err != nil {
		return
	}
}

func (c *UserController) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"error": "Invalid request format"}`, http.StatusBadRequest)
		return
	}

	userPtr, err := c.service.UserService.GetByID(user.ID)
	if err != nil {
		http.Error(w, `{"error": "User not found"}`, http.StatusInternalServerError)
		return
	}
	user = *userPtr

	err = json.NewEncoder(w).Encode(map[string]string{
		"id":       strconv.Itoa(user.ID),
		"username": user.Username,
	})
	if err != nil {
		return
	}
}
