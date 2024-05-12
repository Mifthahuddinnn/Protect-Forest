package user

import (
	"errors"
	"fmt"
	"forest/entities"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) GetUsers() ([]*entities.User, error) {
	var users []*entities.User
	result := repo.DB.Find(&users)
	return users, result.Error
}

func (repo Repository) GetUserByID(id int) (*entities.User, error) {
	var user entities.User
	result := repo.DB.Clauses(clause.Locking{Strength: "UPDATE"}).First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %d not found", id)
		}
		return nil, result.Error
	}
	return &user, nil
}

func (repo Repository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	result := repo.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}

func (repo Repository) DeleteUser(id int) error {
	result := repo.DB.Delete(&entities.User{}, id)
	return result.Error
}

func (repo Repository) LoginUser(email, password string) (*entities.User, error) {
	var user entities.User
	result := repo.DB.Where("email = ?", email).First(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}
	return &user, result.Error
}

func (repo Repository) RegisterUser(user *entities.User) (*entities.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)
	result := repo.DB.Create(user)
	return user, result.Error
}

func (repo Repository) UpdateUser(user *entities.User) error {
	return repo.DB.Save(user).Error
}

func (repo Repository) AddPointsToUser(userID, points int) error {
	user, err := repo.GetUserByID(userID)
	if err != nil {
		return err
	}
	user.Points += points
	return repo.UpdateUser(user)
}
