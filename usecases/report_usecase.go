package usecases

import (
	"forest/entities"
)

type RepoReport interface {
	GetReport() ([]*entities.Reporting, error)
	GetReportByUserID(userID int) ([]*entities.Reporting, error)
	CreateReport(report *entities.Reporting) (*entities.Reporting, error)
	AcceptReport(reportID int, pointsAwarded int) error
}

type ReportUseCase struct {
	Repo RepoReport
}

func (r *ReportUseCase) GetReport() (interface{}, error) {
	return r.Repo.GetReport()
}

func (r *ReportUseCase) GetReportByUserID(userID int) ([]*entities.Reporting, error) {
	return r.Repo.GetReportByUserID(userID)
}

func (r *ReportUseCase) CreateReport(reporting *entities.Reporting) (*entities.Reporting, error) {
	return r.Repo.CreateReport(reporting)
}

func (r *ReportUseCase) AcceptReport(reportID int, pointsAwarded int) error {
	return r.Repo.AcceptReport(reportID, pointsAwarded)
}
