package response

import (
	"forest/entities"
	"time"
)

type ReportResponse struct {
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Photo         string    `json:"photo"`
	ForestAddress string    `json:"forest"`
	Description   string    `json:"description"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}

func FromUseCaseReport(report *entities.Report) *ReportResponse {
	return &ReportResponse{
		Title:         report.Title,
		Content:       report.Content,
		Photo:         report.Photo,
		ForestAddress: report.ForestAddress,
		Description:   report.Description,
		Status:        report.Status,
		CreatedAt:     report.CreatedAt,
	}
}
