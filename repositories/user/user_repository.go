package user

import (
	"errors"
	"fmt"
	"forest/entities"
	"time"

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

func (repo Repository) GetBalanceByUserID(userID int) (*entities.Balance, error) {
	var balance entities.Balance
	result := repo.DB.Where("user_id = ?", userID).First(&balance)
	return &balance, result.Error
}

func (repo Repository) UpdateBalance(balance *entities.Balance) error {
	return repo.DB.Save(balance).Error
}

func (repo Repository) RedeemPoints(userID int) error {
	user, err := repo.GetUserByID(userID)
	if err != nil {
		return err
	}

	if user.Points < 5 {
		return errors.New("Insufficient points for redemption")
	}

	user.Points -= 5
	if err := repo.UpdateUser(user); err != nil {
		return err
	}

	redeem := entities.Redeem{
		UserID:     userID,
		RedeemDate: time.Now(),
	}
	if err := repo.DB.Create(&redeem).Error; err != nil {
		return err
	}

	balance, err := repo.GetBalanceByUserID(userID)
	if err != nil {
		return err
	}

	balance.Amount += 10000
	if err := repo.UpdateBalance(balance); err != nil {
		return err
	}

	return nil
}

func (repo Repository) CreateBalance(balance *entities.Balance) error {
	result := repo.DB.Create(balance)
	return result.Error
}
