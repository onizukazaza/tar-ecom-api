package repository

import (
	"github.com/onizukazaza/tar-ecom-api/entities"
)

type UserRepository interface {
	Listing() ([]*entities.User , error)
	CreateUser(user *entities.User) error
	FindUserByID(id string) (*entities.User, error) 
	EditUser(user *entities.User) error
	//DeleteUserByID
}


