package report

import (
	"forest/entities"
	"gorm.io/gorm"
)

type RepositoryReport struct {
	DB *gorm.DB
}

func (repo RepositoryReport) GetReports() ([]entities.Reporting, error) {
	var reports []entities.Reporting
	result := repo.DB.Find(&reports)
	return reports, result.Error
}

func (repo RepositoryReport) GetReportByID(id int) (*entities.Reporting, error) {
	var report entities.Reporting
	result := repo.DB.Where("id = ?", id).First(&report)
	return &report, result.Error
}

func (repo RepositoryReport) CreateReport(report *entities.Reporting) (*entities.Reporting, error) {
	result := repo.DB.Create(&report)
	return report, result.Error
}

func (repo RepositoryReport) UpdateReport(report *entities.Reporting) (*entities.Reporting, error) {
	result := repo.DB.Save(&report)
	return report, result.Error
}
