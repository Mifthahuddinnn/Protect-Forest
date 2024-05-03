package user

import "forest/entities"

type Userinterface interface {
	GetUsers() ([]entities.User, error)
	GetUserByID(id int) (entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	RegisterUser(user entities.User) (entities.User, error)
	GetUserByEmail(email string) (entities.User, error)
	LoginUser(email, password string) (entities.User, error)
	DeleteUser(id int) error
}
