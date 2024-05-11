package handler

import (
	"forest/entities"
	"forest/handler/base"
	"forest/handler/response"
	"forest/usecases/user"
	"forest/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUseCase user.UserUseCase
}

func (h UserHandler) GetUsers(c echo.Context) error {
	users, err := h.UserUseCase.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get users",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, users)
}

func (h UserHandler) GetUserByID(c echo.Context) error {
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

func (h UserHandler) RegisterUser(c echo.Context) error {
	user := &entities.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}
	createdUser, err := h.UserUseCase.RegisterUser(user)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("User registered successfully", response.FromUseCase(createdUser)))
}

func (h UserHandler) LoginUser(c echo.Context) error {
	loginUser := &entities.User{}
	// TODO mapping di dalam variable
	if err := c.Bind(loginUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request",
			"status":  "400",
		})
	}

	user, err := h.UserUseCase.LoginUser(loginUser.Email, loginUser.Password)
	if err != nil {
		// TODO mapping di dalam variable
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid email or password",
			"status":  "401",
		})
	}

	token, err := utils.CreateToken(user.ID)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewLoginResponse("Login success", token))
}
