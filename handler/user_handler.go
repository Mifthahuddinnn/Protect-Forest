package handler

import (
	"forest/entities"
	"forest/usecases"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type UserHandler struct {
	UserUseCase usecases.UserUseCase
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.UserUseCase.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get users",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid id",
			"error":   err.Error(),
		})
	}
	user, err := h.UserUseCase.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get user",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	user := &entities.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request",
			"status":  "400",
		})
	}

	createdUser, err := h.UserUseCase.RegisterUser(*user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create user",
			"status":  "500",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "User created successfully",
		"status":  "200",
		"data":    createdUser,
	})

}

func (h *UserHandler) LoginUser(c echo.Context) error {
	loginUser := &entities.User{}
	if err := c.Bind(loginUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request",
			"status":  "400",
		})
	}

	user, err := h.UserUseCase.LoginUser(loginUser.Email, loginUser.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid email or password",
			"status":  "401",
		})
	}

	token, err := createToken(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create token",
			"status":  "500",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Logged in successfully",
		"status":  "200",
		"data":    token,
	})
}

func createToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}
