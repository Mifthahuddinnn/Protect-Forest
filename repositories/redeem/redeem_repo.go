package redeem

import (
	"forest/entities"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetRedeems() ([]*entities.Redeem, error) {
	var redeems []*entities.Redeem
	result := repo.DB.Find(&redeems)
	return redeems, result.Error
}


