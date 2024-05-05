package utils

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type JwtToken struct {
}

func (JwtToken) CreateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}
