package handler

import (
	"forest/entities"
	"forest/usecases"
	"forest/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AdminHandler struct {
	AdminUseCase usecases.AdminUseCase
}

func (ah AdminHandler) RegisterAdmin(c echo.Context) error {
	admin := &entities.Admin{}
	if err := c.Bind(admin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}
	registeredAdmin, err := ah.AdminUseCase.RegisterAdmin(admin)
	if err != nil {
		if err.Error() == "username already exist" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Username already exist",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to register admin",
			"error":   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, registeredAdmin)
}

func (ah AdminHandler) LoginAdmin(c echo.Context) error {
	loginAdmin := &entities.Admin{}
	if err := c.Bind(loginAdmin); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}
	admin, err := ah.AdminUseCase.LoginAdmin(loginAdmin.Username, loginAdmin.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid username or password",
			"error":   err.Error(),
		})

	}
	token, err := utils.CreateToken(admin.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create token",
			"error":   err.Error(),
		})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login success",
		"token":   token,
		"data":    admin,
	})
}
