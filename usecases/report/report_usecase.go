package report

import (
	"fmt"
	"forest/entities"
	"forest/usecases/user"
)

type Repository interface {
	CreateReport(report *entities.Report) (*entities.Report, error)
	GetReports() ([]*entities.Report, error)
	GetReportByID(id int) (*entities.Report, error)
	UpdateReport(Title string, Content string, Status string) (*entities.Report, error)
	DeleteReport(id int) error
}

type ReportUseCase struct {
	Repo     Repository
	UserRepo user.Repository
}

func (r ReportUseCase) CreateReport(report *entities.Report) (*entities.Report, error) {
	_, err := r.UserRepo.GetUserByID(report.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return r.Repo.CreateReport(report)
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
