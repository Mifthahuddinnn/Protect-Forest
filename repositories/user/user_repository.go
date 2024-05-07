package user

import (
	"forest/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetUsers() ([]*UserDB, error) {
	var users []*UserDB
	result := repo.DB.Find(&users)
	return users, result.Error
}

func (repo Repository) GetUserByID(id int) (*UserDB, error) {
	var user UserDB
	result := repo.DB.Where("id = ?", id).First(&user)
	return &user, result.Error
}

func (repo Repository) UpdateUser(user *UserDB) (*UserDB, error) {
	result := repo.DB.Model(&entities.User{}).Where("id = ?", user.ID).Updates(user)
	return user, result.Error
}

func (repo Repository) GetUserByEmail(email string) (*UserDB, error) {
	var user UserDB
	result := repo.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (repo Repository) DeleteUser(id int) error {
	result := repo.DB.Delete(&UserDB{}, id)
	return result.Error
}

func (repo Repository) LoginUser(email, password string) (*UserDB, error) {
	var user UserDB
	result := repo.DB.Where("email = ?", email).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return &user, result.Error
}

func (repo Repository) RegisterUser(user *UserDB) (*UserDB, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	result := repo.DB.Create(user)
	return user, result.Error
}
