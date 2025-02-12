package service

import (
	"errors"
	"github.com/Crabocod/gpt_network/api-service/internal/app/store"
	"github.com/Crabocod/gpt_network/api-service/internal/models"
	"github.com/Crabocod/gpt_network/api-service/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

type UserServiceInterface interface {
	Save(user models.User) error
	Get(username, passwordHash string) (*models.User, error)
	GetIDByToken(refreshToken string) (int, error)
	GetByID(id int) (*models.User, error)
	GetByName(name string) (*models.User, error)
}

type UserService struct {
	store store.Store
}

func NewUserService(s store.Store) UserServiceInterface {
	return &UserService{
		store: s,
	}
}

func (s *UserService) Save(user models.User) error {
	user.PasswordHash = utils.HashPassword(user.Password)
	err := s.store.User().Save(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Get(username, password string) (*models.User, error) {
	passwordHash := utils.HashPassword(password)
	user, err := s.store.User().Get(username, passwordHash)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetByID(id int) (*models.User, error) {
	user, err := s.store.User().GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetByName(name string) (*models.User, error) {
	user, err := s.store.User().GetUserByName(name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetIDByToken(refreshToken string) (int, error) {
	claims := &utils.JWTClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return utils.RefreshSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("token is not valid")
	}

	return claims.UserID, nil
}
