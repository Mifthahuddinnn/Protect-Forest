package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.ErrUnauthorized
		}

		tokenString := strings.SplitN(authHeader, " ", 2)[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ambil nilai variabel lingkungan JWT_SECRET dari ENV_PROJECT
			environment := os.Getenv("ENV_PROJECT")
			envs := strings.Split(environment, " ")

			var envMap map[string]string
			envMap = make(map[string]string)
			for _, env := range envs {
				keyValue := strings.Split(env, "=")
				envMap[keyValue[0]] = keyValue[1]
			}

			return []byte(envMap["JWT_SECRET"]), nil
		})
		if err != nil {
			fmt.Println("Error parsing token:", err)
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				return echo.ErrUnauthorized
			}
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
		}

		if !token.Valid {
			fmt.Println("Token is invalid")
			return echo.ErrUnauthorized
		}

		claims := token.Claims.(jwt.MapClaims)
		userId := claims["user_id"].(float64)
		c.Set("user_id", int(userId))

		return next(c)
	}
}
