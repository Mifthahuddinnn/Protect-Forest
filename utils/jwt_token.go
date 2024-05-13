package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
}

func CreateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret"
	}
	return token.SignedString([]byte(secret))
}
