package response

import (
	"forest/entities"
	"time"
)

type ReportResponse struct {
	ID            int
	Title         string
	Content       string
	Photo         string
	ForestAddress string
	Description   string
	Status        string
	CreatedAt     time.Time
}

func FromUseCaseReport(report *entities.Report) *ReportResponse {
	return &ReportResponse{
		ID:            report.ID,
		Title:         report.Title,
		Content:       report.Content,
		Photo:         report.Photo,
		ForestAddress: report.ForestAddress,
		Description:   report.Description,
		Status:        report.Status,
		CreatedAt:     report.CreatedAt,
	}
}
