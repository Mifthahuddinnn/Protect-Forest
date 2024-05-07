package report

import (
	"gorm.io/gorm"
)

type RepositoryReport struct {
	DB *gorm.DB
}

func (repo RepositoryReport) GetReports() ([]ReportingDB, error) {
	var reports []ReportingDB
	result := repo.DB.Find(&reports)
	return reports, result.Error
}

func (repo RepositoryReport) GetReportByID(id int) (*ReportingDB, error) {
	var report ReportingDB
	result := repo.DB.Where("id = ?", id).First(&report)
	return &report, result.Error
}

func (repo RepositoryReport) CreateReport(report *ReportingDB) (*ReportingDB, error) {
	result := repo.DB.Create(&report)
	return report, result.Error
}

func (repo RepositoryReport) UpdateReport(report *ReportingDB) (*ReportingDB, error) {
	result := repo.DB.Save(&report)
	return report, result.Error
}
