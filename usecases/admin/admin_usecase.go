package admin

import (
	"errors"
	"forest/constant"
	"forest/entities"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type RepoAdmin interface {
	RegisterAdmin(admin *entities.Admin) (*entities.Admin, error)
	LoginAdmin(username, password string) (*entities.Admin, error)
	GetAdminUsername(username string) (*entities.Admin, error)
}

type UseCaseAdmin struct {
	Repo RepoAdmin
}

func (a UseCaseAdmin) RegisterAdmin(admin *entities.Admin) (*entities.Admin, error) {
	if admin.Username == "" || admin.Password == "" {
		return nil, constant.ErrorAdminEmptyField
	}
	existingUsername, err := a.Repo.GetAdminUsername(admin.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUsername.Username != "" {
		return nil, constant.ErrorUsernameAlreadyExist
	}
	return a.Repo.RegisterAdmin(admin)
}

func (a UseCaseAdmin) LoginAdmin(username, password string) (*entities.Admin, error) {
	if username == "" || password == "" {
		return nil, constant.ErrorAdminEmptyField
	}
	existingAdmin, err := a.Repo.GetAdminUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constant.ErrorUsernameNotFound
		}
		return nil, err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(existingAdmin.Password), []byte(password)); err != nil {
		return nil, constant.ErrorIncorrectPassword
	}
	return a.Repo.LoginAdmin(username, password)
}
