package user

import (
	"encoding/json"
	"errors"
	"forest/constant"
	"forest/entities"
	"io/ioutil"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	RegisterUser(*entities.User) (*entities.User, error)
	LoginUser(email, password string) (*entities.User, error)
	GetUserByID(id int) (*entities.User, error)
	GetUsers() ([]*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(*entities.User) error
	AddPointsToUser(userID, points int) error
	RedeemPoints(userID int) error
	GetBalanceByUserID(userID int) (*entities.Balance, error)
	UpdateBalance(*entities.Balance) error
	CreateBalance(*entities.Balance) error
}

type UserUseCase struct {
	Repo Repository
}

func (u *UserUseCase) RegisterUser(user *entities.User) (*entities.User, error) {
	if user.Email == "" || user.Password == "" || user.Name == "" {
		return nil, constant.ErrorEmptyInput
	}

	if !strings.Contains(user.Email, "@") {
		return nil, constant.ErrorEmailInvalid
	}

	if len(user.Password) < 6 {
		return nil, constant.ErrorPassword
	}

	existingUser, err := u.Repo.GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingUser.Email != "" {
		return nil, constant.ErrorEmailExists
	}

	// Register user
	registeredUser, err := u.Repo.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	// Create default balance record
	defaultBalance := &entities.Balance{
		UserID: registeredUser.ID,
		Amount: 0, // Set initial balance amount here if needed
	}
	if err := u.Repo.CreateBalance(defaultBalance); err != nil {
		return nil, err
	}

	return registeredUser, nil
}

func (u *UserUseCase) LoginUser(email, password string) (*entities.User, error) {
	return u.Repo.LoginUser(email, password)
}

func (u *UserUseCase) GetUserByID(id int) (*entities.User, error) {
	return u.Repo.GetUserByID(id)
}

func (u *UserUseCase) GetUsers() ([]*entities.User, error) {
	return u.Repo.GetUsers()
}

func (u *UserUseCase) AddPointsToUser(userID, points int) error {
	return u.Repo.AddPointsToUser(userID, points)
}

func (u *UserUseCase) RedeemPoints(userID int) error {
	return u.Repo.RedeemPoints(userID)
}

func (u *UserUseCase) GetBalanceByUserID(userID int) (*entities.Balance, error) {
	return u.Repo.GetBalanceByUserID(userID)
}

func (u *UserUseCase) UpdateBalance(balance *entities.Balance) error {
	return u.Repo.UpdateBalance(balance)
}

func (u *UserUseCase) GetNews() (*entities.NewsResponse, error) {
	resp, err := http.Get("https://newsdata.io/api/1/news?apikey=pub_441559a0bbb34983f12e0cf2d0f52329f4f2c&q=hutan&country=id&language=id")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var newsResponse entities.NewsResponse
	err = json.Unmarshal(bodyBytes, &newsResponse)
	if err != nil {
		return nil, err
	}

	return &newsResponse, nil
}
