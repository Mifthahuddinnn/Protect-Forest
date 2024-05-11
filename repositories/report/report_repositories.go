package report

import (
	"forest/entities"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetReports() ([]*entities.Report, error) {
	var reports []*entities.Report
	result := repo.DB.Find(&reports)
	return reports, result.Error
}

func (repo Repository) GetReportByID(id int) (*entities.Report, error) {
	var report entities.Report
	result := repo.DB.Where("id = ?", id).First(&report)
	return &report, result.Error
}

func (repo Repository) CreateReport(report *entities.Report) (*entities.Report, error) {
	result := repo.DB.Create(report)
	if result.Error != nil {
		return nil, result.Error
	}
	return report, nil
}

func (repo Repository) UpdateReport(Title string, Content string, Status string) (*entities.Report, error) {
	report := &entities.Report{
		Title:   Title,
		Content: Content,
		Status:  Status,
	}
	result := repo.DB.Save(report)
	return report, result.Error
}

func (repo Repository) DeleteReport(id int) error {
	result := repo.DB.Delete(&entities.Report{}, id)
	return result.Error
}
