package repository

import (
	"github.com/jmoiron/sqlx"
	// "github.com/gofiber/fiber/v2/log"
	// "github.com/onizukazaza/tar-ecom-api/entities"
	// _userlistexception "github.com/onizukazaza/tar-ecom-api/pkg/admin/exception"
)

type adminRepositoryImpl struct {
	db     *sqlx.DB
}


// func NewAdminRepositoryImpl(db *sqlx.DB ) *adminRepositoryImpl {
// 	return &adminRepositoryImpl{db: db}
// }


// func (r *adminRepositoryImpl) Listing() ([]*entities.User, error) {
// 	userList := make([]*entities.User, 0)

// 	query := "SELECT * FROM users"

// 	err := r.db.Select(&userList, query)
// 	if err != nil {
// 		log.Errorf("Failed to list users: %v", err)
// 		return nil, &_userlistexception.UserListing{}
// 	}

// 	return userList, nil
// }
