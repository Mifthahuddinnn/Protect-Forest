package repositories

import (
	"forest/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo UserRepository) GetUsers() ([]*entities.User, error) {
	var users []*entities.User
	result := repo.DB.Find(&users)
	return users, result.Error
}

func (repo UserRepository) GetUserByID(id int) (*entities.User, error) {
	var user entities.User
	result := repo.DB.Where("id = ?", id).First(&user)
	return &user, result.Error
}

func (repo UserRepository) UpdateUser(user *entities.User) (*entities.User, error) {
	result := repo.DB.Model(&entities.User{}).Where("id = ?", user.ID).Updates(user)
	return user, result.Error
}

func (repo UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	result := repo.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (repo UserRepository) DeleteUser(id int) error {
	result := repo.DB.Delete(&entities.User{}, id)
	return result.Error
}

func (repo UserRepository) LoginUser(email, password string) (*entities.User, error) {
	var user entities.User
	result := repo.DB.Where("email = ?", email).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return &user, result.Error
}

func (repo UserRepository) RegisterUser(user *entities.User) (*entities.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	result := repo.DB.Create(user)
	return user, result.Error
}
