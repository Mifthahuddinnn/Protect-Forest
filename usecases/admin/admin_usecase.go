package admin

import (
	"errors"
	"forest/entities"
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
		return nil, errors.New("username and password are required")
	}
	existingUsername, err := a.Repo.GetAdminUsername(admin.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existingUsername.Username != "" {
		return nil, errors.New("username already exist")
	}
	return a.Repo.RegisterAdmin(admin)
}

func (a UseCaseAdmin) LoginAdmin(username, password string) (*entities.Admin, error) {
	return a.Repo.LoginAdmin(username, password)
}
