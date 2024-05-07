package admin

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RepositoryAdmin struct {
	DB *gorm.DB
}

func (repo RepositoryAdmin) RegisterAdmin(admin *DB) (*DB, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	admin.Password = string(hashedPassword)
	result := repo.DB.Create(&admin)
	return admin, result.Error
}

func (repo RepositoryAdmin) LoginAdmin(username, password string) (*DB, error) {
	var admin DB
	result := repo.DB.Where("username = ?", username).First(&admin)
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return nil, err
	}
	return &admin, result.Error
}

func (repo RepositoryAdmin) GetAdminUsername(username string) (*DB, error) {
	var admin DB
	result := repo.DB.Where("username = ?", username).First(&admin)
	return &admin, result.Error
}
