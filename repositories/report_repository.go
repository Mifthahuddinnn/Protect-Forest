package repositories

import (
	"forest/entities"
	"gorm.io/gorm"
)

type RepoReport struct {
	DB             *gorm.DB
	UserRepository *UserRepository
}

func (r *RepoReport) GetReport() ([]*entities.Reporting, error) {
	var reports []*entities.Reporting
	err := r.DB.Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (r *RepoReport) GetReportByUserID(userID int) ([]*entities.Reporting, error) {
	var reports []*entities.Reporting
	err := r.DB.Where("id = ?", userID).Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (r *RepoReport) CreateReport(report *entities.Reporting) (*entities.Reporting, error) {
	err := r.DB.Create(report).Error
	if err != nil {
		return nil, err
	}
	return report, nil
}

//func (r *RepoReport) AcceptReport(reportID int, pointsAwarded int) error {
//	var report entities.Reporting
//	err := r.DB.Model(&entities.Reporting{}).Where("id = ?", reportID).Update("status", "accepted").First(&report).Error
//	if err != nil {
//		return err
//	}
//
//	return r.UserRepository.UpdateUserPoints(report.UserID, pointsAwarded)
//}
