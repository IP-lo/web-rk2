package usecase

import "github.com/IP-lo/web-rk2/internal/entities"

type Provider interface {
	InsertUser(entities.User) (*entities.User, error)
	SelectAllUsers() ([]*entities.User, error)

	SelectUserByID(id int) (*entities.User, error)
	SelectUserByName(name string) (*entities.User, error)
	SelectUserByEmail(name string) (*entities.User, error)

	UpdateUserByID(id int, user entities.User) (*entities.User, error)
	DeleteUserByID(id int) error
}
