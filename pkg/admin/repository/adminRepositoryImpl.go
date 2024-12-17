package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/onizukazaza/tar-ecom-api/entities"
)

type adminRepositoryImpl struct {
	db     *sqlx.DB
	logger *log.Logger
}


func NewAdminRepositoryImpl(db *sqlx.DB, logger *log.Logger) AdminRepository {
	return &adminRepositoryImpl{
		db:     db,
		logger: logger,
	}
}


func (r *adminRepositoryImpl) Listing() ([]*entities.User, error) {
	userList := make([]*entities.User, 0)

	query := "SELECT * FROM users"

	err := r.db.Select(&userList, query)
	if err != nil {
		r.logger.Printf("Failed to list users: %v", err)
		return nil, err
	}

	return userList, nil
}
