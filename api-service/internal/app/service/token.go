package service

import (
	"github.com/Crabocod/gpt_network/api-service/internal/app/store"
	"github.com/Crabocod/gpt_network/api-service/internal/utils"
)

type TokenServiceInterface interface {
	Save(userID int, refreshToken string) error
	GenerateAccess(userID int) (string, error)
	GenerateRefresh(userID int) (string, error)
	GetByUserID(userID int) (string, error)
	Delete(userID int) error
}

type TokenService struct {
	store store.Store
}

func NewTokenService(s store.Store) *TokenService {
	return &TokenService{
		store: s,
	}
}

func (s *TokenService) Save(userID int, refreshToken string) error {
	err := s.store.Token().Save(userID, refreshToken)
	if err != nil {
		return err
	}

	return nil
}

func (s *TokenService) Delete(userID int) error {
	err := s.store.Token().Delete(userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *TokenService) GenerateAccess(userID int) (string, error) {
	return utils.GenerateJWT(userID)
}

func (s *TokenService) GenerateRefresh(userID int) (string, error) {
	return utils.GenerateRefreshToken(userID)
}

func (s *TokenService) GetByUserID(userID int) (string, error) {
	refreshToken, err := s.store.Token().GetByUserID(userID)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
