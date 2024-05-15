package handler

import (
	"context"
	"forest/config"
	"forest/entities"
	"forest/handler/base"
	"forest/handler/response"
	"forest/usecases/report"
	"forest/utils"
	"net/http"
	"strconv"

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
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, base.NewSuccessResponse("Report created successfully", response.FromUseCaseReport(createdReport)))
}

func (h *ReportHandler) ApproveReport(c echo.Context) error {
	reportIDStr := c.Param("id")
	reportID, err := strconv.Atoi(reportIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid report ID",
		})
	}

	adminID := c.Get("user_id").(int)
	var approvedReport *entities.Report

	approvedReport, err = h.ReportUseCase.ApproveReport(reportID, adminID)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Report approved successfully", response.FromUseCaseApprove(approvedReport)))
}

func (h *ReportHandler) DeleteReport(c echo.Context) error {
	reportIDStr := c.Param("id")
	reportID, err := strconv.Atoi(reportIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid report ID",
		})
	}

	err = h.ReportUseCase.DeleteReport(reportID)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewDeleteBase("Report deleted successfully"))

}

func (h *ReportHandler) GetReports(c echo.Context) error {
	reports, err := h.ReportUseCase.GetReports()
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, base.NewSuccessResponse("Reports retrieved successfully", response.FromGetReports(reports)))
}

func (h *ReportHandler) GetReportByID(c echo.Context) error {
	reportIDStr := c.Param("id")
	reportID, err := strconv.Atoi(reportIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Invalid report ID",
		})
	}

	report, err := h.ReportUseCase.GetReportByID(reportID)
	if err != nil {
		return c.JSON(utils.ConvertResponseCode(err), base.NewErrResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, base.NewSuccessResponse("Report retrieved successfully", response.FromUseCaseReport(report)))
}
