package repositories

import (
	"forest/entities"
	"gorm.io/gorm"
)

type RedeemRepository struct {
	DB *gorm.DB
}

func (repo RedeemRepository) GetRedeems() ([]*entities.Redeem, error) {
	var redeems []*entities.Redeem
	result := repo.DB.Find(&redeems)
	return redeems, result.Error
}

func (repo RedeemRepository) GetRedeemByID(id int) (*entities.Redeem, error) {
	var redeem entities.Redeem
	result := repo.DB.Where("id = ?", id).First(&redeem)
	return &redeem, result.Error
}
