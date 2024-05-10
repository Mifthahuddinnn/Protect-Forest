package handler

import (
	"context"
	"forest/config"
	"forest/entities"
	"forest/handler/base"
	"forest/usecases/report"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ReportHandler struct {
	ReportUseCase report.ReportUseCase
}

func (h ReportHandler) CreateReport(c echo.Context) error {
	fileHeader, err := c.FormFile("photo")
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrResponse("Failed to get file from form"))
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrResponse("Failed to open file"))
	}
	defer file.Close()

	cloudinaryService, err := config.InitCloudinary()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrResponse("Failed to initialize Cloudinary service"))
	}
	uploadResult, err := cloudinaryService.Upload.Upload(context.Background(), file, uploader.UploadParams{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrResponse("Failed to upload file to Cloudinary"))
	}

	report := &entities.Report{}
	if err := c.Bind(report); err != nil {
		return c.JSON(http.StatusBadRequest, base.NewErrResponse("Invalid request"))
	}

	report.Photo = uploadResult.SecureURL

	createdReport, err := h.ReportUseCase.CreateReport(report.Title, report.Content, report.ForestAddress, report.Description, report.Photo, report.Status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, base.NewErrResponse("Failed to create report"))
	}

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Report created successfully", createdReport))
}
