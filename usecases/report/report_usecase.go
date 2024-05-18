package report

import (
	"fmt"
	"forest/constant"
	"forest/entities"
	"forest/usecases/user"
)

type Repository interface {
	CreateReport(report *entities.Report) (*entities.Report, error)
	GetReports() ([]*entities.Report, error)
	GetReportByID(id int) (*entities.Report, error)
	DeleteReport(id int) error
	UpdateReport(report *entities.Report) (*entities.Report, error)
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
	if report.Title == "" {
		return nil, constant.ErrorTitleEmpty
	}
	if report.Content == "" {
		return nil, constant.ErrorContentEmpty
	}

	return r.Repo.CreateReport(report)
}

func (r ReportUseCase) GetReports() ([]*entities.Report, error) {
	return r.Repo.GetReports()
}

func (r ReportUseCase) GetReportByID(id int) (*entities.Report, error) {
	return r.Repo.GetReportByID(id)
}

func (r ReportUseCase) DeleteReport(id int) error {
	report, err := r.Repo.GetReportByID(id)
	if err != nil {
		return err
	}
	if report == nil {
		return fmt.Errorf("Report with ID %d not found", id)
	}
	if report.Status == "approved" {
		return constant.ErrorReportAlreadyApproved
	}
	return r.Repo.DeleteReport(id)
}

func (r ReportUseCase) ApproveReport(reportID int, adminID int) (*entities.Report, error) {
	report, err := r.Repo.GetReportByID(reportID)
	if err != nil {
		return nil, constant.ErrorReportNotFound
	}

	if report.Status == "approved" {
		return nil, constant.ErrorReportAlreadyApproved
	}

	report.Status = "approved"
	report.AdminID = &adminID

	updatedReport, err := r.Repo.UpdateReport(report)
	if err != nil {
		return nil, fmt.Errorf("failed to approve report: %w", err)
	}

	err = r.UserRepo.AddPointsToUser(report.UserID, 1)
	if err != nil {
		return nil, fmt.Errorf("failed to add points to user: %w", err)
	}

	return updatedReport, nil
}
