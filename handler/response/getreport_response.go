package response

import "forest/entities"

type GetReportResponse struct {
	ID            int
	Username      string
	Title         string
	Content       string
	Photo         string
	ForestAddress string
	Description   string
	Status        string
	CreatedAt     string
}

func FromGetReport(report *entities.Report) *GetReportResponse {
	return &GetReportResponse{
		ID:            report.ID,
		Username:      report.User.Name,
		Title:         report.Title,
		Content:       report.Content,
		Photo:         report.Photo,
		ForestAddress: report.ForestAddress,
		Description:   report.Description,
		Status:        report.Status,
		CreatedAt:     report.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func FromGetReports(reports []*entities.Report) []*GetReportResponse {
	var reportsResponse []*GetReportResponse
	for _, report := range reports {
		reportsResponse = append(reportsResponse, FromGetReport(report))
	}
	return reportsResponse
}
