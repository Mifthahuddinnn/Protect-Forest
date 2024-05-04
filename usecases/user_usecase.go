package usecases

import (
	"errors"
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
	if user.Email == "" || user.Password == "" {
		return nil, errors.New("email and password are required")
	}

	if !strings.Contains(user.Email, "@") {
		return nil, errors.New("email is invalid")
	}

	if len(user.Password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}

	existingUser, err := u.Repo.GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingUser.Email != "" {
		return nil, errors.New("email already exists")
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
