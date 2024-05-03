package usecases

import (
	"errors"
	"forest/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (entities.User, error)
	RegisterUser(user entities.User) (entities.User, error)
	LoginUser(email, password string) (entities.User, error)
	GetUserByID(id int) (entities.User, error)
	GetUsers() ([]entities.User, error)
}

type UserUseCase struct {
	UserRepository UserRepository
}

func (u *UserUseCase) RegisterUser(user entities.User) (entities.User, error) {
	if user.Email == "" || user.Password == "" {
		return entities.User{}, errors.New("email and password are required")
	}

	existingUser, err := u.UserRepository.GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return entities.User{}, err
	}

	if existingUser.Email != "" {
		return entities.User{}, errors.New("user already exists")
	}

	return u.UserRepository.RegisterUser(user)
}

func (u *UserUseCase) LoginUser(email, password string) (entities.User, error) {
	return u.UserRepository.LoginUser(email, password)
}

func (u *UserUseCase) GetUserByID(id int) (entities.User, error) { // New method
	return u.UserRepository.GetUserByID(id)
}

func (u *UserUseCase) GetUsers() ([]entities.User, error) { // New method
	return u.UserRepository.GetUsers()
}
