package user

import (
	"encoding/json"
	"errors"
	"forest/constant"
	"forest/entities"
	"io/ioutil"
	"net/http"
	"os"
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

	registeredUser, err := u.Repo.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	defaultBalance := &entities.Balance{
		UserID: registeredUser.ID,
		Amount: 0,
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
	user, err := u.GetUserByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return constant.ErrorUserNotFound
	}
	if user.Points < 5 {
		return constant.ErrorPointNotEnough
	}

	// Subtract points from the user
	user.Points -= 5

	// Update the user in the repository
	err = u.Repo.UpdateUser(user)
	if err != nil {
		return err
	}

	// Get the user's balance
	balance, err := u.Repo.GetBalanceByUserID(userID)
	if err != nil {
		return err
	}

	// Add to the user's balance
	balance.Amount += 10000

	// Update the balance in the repository
	err = u.Repo.UpdateBalance(balance)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) GetBalanceByUserID(userID int) (*entities.Balance, error) {
	return u.Repo.GetBalanceByUserID(userID)
}

func (u *UserUseCase) UpdateBalance(balance *entities.Balance) error {
	return u.Repo.UpdateBalance(balance)
}

func (u *UserUseCase) UpdateUser(user *entities.User) error {
	return u.Repo.UpdateUser(user)
}

func (u *UserUseCase) GetNews() (*entities.NewsResponse, error) {
	apiKey := os.Getenv("NEWS_API_KEY")
	resp, err := http.Get("https://newsdata.io/api/1/news?apikey=" + apiKey + "&q=hutan&country=id&language=id")
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
