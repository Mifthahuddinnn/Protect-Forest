package handler

import (
	"forest/entities"
	"forest/handler/base"
	"forest/usecases/report"
	"forest/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ReportHandler struct {
	ReportUseCase report.UseCaseReport
}

func (h ReportHandler) CreateReport(c echo.Context) error {
	report := &entities.Reporting{}
	if err := c.Bind(report); err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}
	createdReport, err := h.ReportUseCase.CreateReport(report)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, createdReport)
}
