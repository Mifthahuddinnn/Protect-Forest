package report

import "forest/entities"

type Repository interface {
	CreateReport(Title string, Content string, Photo string, Status string, ForestAddress string, Description string) (*entities.Report, error)
	GetReports() ([]*entities.Report, error)
	GetReportByID(id int) (*entities.Report, error)
	UpdateReport(Title string, Content string, Status string) (*entities.Report, error)
	DeleteReport(id int) error
}

type ReportUseCase struct {
	Repo Repository
}

func (r ReportUseCase) CreateReport(Title string, Content string, Photo string, Status string, ForestAddress string, Description string) (*entities.Report, error) {
	return r.Repo.CreateReport(Title, Content, Photo, Status, ForestAddress, Description)
}

func (r ReportUseCase) GetReports() ([]*entities.Report, error) {
	return r.Repo.GetReports()
}

func (r ReportUseCase) GetReportByID(id int) (*entities.Report, error) {
	return r.Repo.GetReportByID(id)
}

func (r ReportUseCase) UpdateReport(Title string, Content string, Status string) (*entities.Report, error) {
	return r.Repo.UpdateReport(Title, Content, Status)
}

func (r ReportUseCase) DeleteReport(id int) error {
	return r.Repo.DeleteReport(id)
}
