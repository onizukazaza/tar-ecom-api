package repository

import (
	"github.com/onizukazaza/tar-ecom-api/entities"
)


type AdminRepository interface {
	Listing() ([]*entities.User , error) 

}

