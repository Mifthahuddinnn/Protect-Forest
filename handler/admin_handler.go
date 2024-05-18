package handler

import (
	"forest/entities"
	"forest/handler/base"
	"forest/handler/response"
	"forest/usecases/admin"
	"forest/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	AdminUseCase admin.UseCaseAdmin
}

func (ah AdminHandler) RegisterAdmin(c echo.Context) error {
	admin := &entities.Admin{}
	if err := c.Bind(admin); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}
	registeredAdmin, err := ah.AdminUseCase.RegisterAdmin(admin)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Admin registered successfully", response.FromAdmin(registeredAdmin)))
}

func (ah AdminHandler) LoginAdmin(c echo.Context) error {
	loginAdmin := &entities.Admin{}
	if err := c.Bind(loginAdmin); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}
	admin, err := ah.AdminUseCase.LoginAdmin(loginAdmin.Username, loginAdmin.Password)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))

	}
	token, err := utils.CreateToken(admin.ID)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewLoginResponse("Login Success", token))
}
