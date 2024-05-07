package report

import "forest/entities"

type ReportRepo interface {
	GetReports() ([]entities.Reporting, error)
	GetReportByID(id int) (*entities.Reporting, error)
	CreateReport(report *entities.Reporting) (*entities.Reporting, error)
}

type UseCaseReport struct {
	Repo ReportRepo
}

func (u *UseCaseReport) GetReports() ([]entities.Reporting, error) {
	return u.Repo.GetReports()
}

func (u *UseCaseReport) GetReportByID(id int) (*entities.Reporting, error) {
	return u.Repo.GetReportByID(id)
}

func (u *UseCaseReport) CreateReport(report *entities.Reporting) (*entities.Reporting, error) {
	return u.Repo.CreateReport(report)
}
