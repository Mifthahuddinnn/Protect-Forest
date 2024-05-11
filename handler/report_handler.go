package handler

import (
	"context"
	"forest/config"
	"forest/entities"
	"forest/handler/base"
	"forest/handler/response"
	"forest/usecases/report"
	"net/http"

	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/labstack/echo/v4"
)

type ReportHandler struct {
	ReportUseCase report.ReportUseCase
}

func (h *ReportHandler) CreateReport(c echo.Context) error {
	fileHeader, err := c.FormFile("photo")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Failed to get file from form",
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Failed to open file",
		})
	}
	defer file.Close()

	cloudinaryService, err := config.InitCloudinary()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to initialize Cloudinary service",
		})
	}
	uploadResult, err := cloudinaryService.Upload.Upload(context.Background(), file, uploader.UploadParams{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to upload file to Cloudinary",
		})
	}

	report := new(entities.Report)
	if err := c.Bind(report); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid request",
		})
	}
	report.Photo = uploadResult.SecureURL

	userID, ok := c.Get("user_id").(int)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "Invalid user token",
		})
	}
	report.UserID = userID

	createdReport, err := h.ReportUseCase.CreateReport(report)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   "Failed to create report",
			"details": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Report created successfully", response.FromUseCaseReport(createdReport)))
}
