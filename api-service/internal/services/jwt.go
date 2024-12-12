package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"web.app/internal/config"
)

var JWTSecret = []byte(config.Data.JWT.AccessSecret)
var RefreshSecret = []byte(config.Data.JWT.RefreshSecret)

type JWTClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

func GenerateRefreshToken(userID int) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(RefreshSecret)
}
