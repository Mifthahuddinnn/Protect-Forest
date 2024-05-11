package utils

import (
	"github.com/golang-jwt/jwt"
	"os"
	"time"
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
