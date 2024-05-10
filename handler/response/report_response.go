package response

import "forest/entities"

type SuccessResponse struct {
	Message string
	Status  bool
	Data    interface{}
}

func FromUseCaseReport(report *entities.Report) *entities.Report {
	return &entities.Report{
		Title:     report.Title,
		Content:   report.Content,
		Photo:     report.Photo,
		Status:    report.Status,
		CreatedAt: report.CreatedAt,
	}
}
