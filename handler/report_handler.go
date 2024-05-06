package handler

import (
	"context"
	"forest/config"
	"forest/entities"
	"forest/handler/base"
	"forest/usecases"
	"forest/utils"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"net/http"
	"strconv"
)

type ReportHandler struct {
	ReportUseCase usecases.ReportUseCase
}

func (rh ReportHandler) GetReport(c echo.Context) error {
	report, err := rh.ReportUseCase.GetReport()
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Report fetched successfully", report))
}

func (rh ReportHandler) GetReportByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	report, err := rh.ReportUseCase.GetReportByUserID(id)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Report fetched successfully", report))
}

func (rh ReportHandler) CreateReport(c echo.Context) error {
	fileHeader, err := c.FormFile("Photo")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to get file from form",
			"error":   err.Error(),
		})
	}

	file, err := fileHeader.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed to open file",
			"error":   err.Error(),
		})
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	cloudinaryService, err := config.InitCloudinary()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to initialize Cloudinary service",
			"error":   err.Error(),
		})
	}
	uploadResult, err := cloudinaryService.Upload.Upload(context.Background(), file, uploader.UploadParams{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to upload file to Cloudinary",
			"error":   err.Error(),
		})
	}

	report := &entities.Reporting{}
	if err := c.Bind(report); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request",
		})
	}

	report.Photo = uploadResult.SecureURL

	createdReport, err := rh.ReportUseCase.CreateReport(report)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create report data",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Report created successfully", createdReport))

}
