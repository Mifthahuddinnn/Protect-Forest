package usecases

import (
	"errors"
	"forest/constant"
	"forest/entities"
	"gorm.io/gorm"
	"strings"
)

type Repository interface {
	RegisterUser(*entities.User) (*entities.User, error)
	LoginUser(email, password string) (*entities.User, error)
	GetUserByID(id int) (*entities.User, error)
	GetUsers() ([]*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
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

	return u.Repo.RegisterUser(user)
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
