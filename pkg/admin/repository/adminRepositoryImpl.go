package repository

import (
	"github.com/jmoiron/sqlx"
	_AdminException "github.com/onizukazaza/tar-ecom-api/pkg/admin/exception"
	"fmt"
)

type adminRepositoryImpl struct {
	db *sqlx.DB
}

func NewAdminRepositoryImpl(db *sqlx.DB) AdminRepository {
	return &adminRepositoryImpl{db: db}
}

func (r *adminRepositoryImpl) UpdateUserRole(userID string, role string) error {
	query := `
        UPDATE users 
        SET role = $1 
        WHERE id = $2
    `
	result, err := r.db.Exec(query, role, userID)
	if err != nil {
		// SQL Error
		return fmt.Errorf("UpdateUserRole: failed to execute query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
	
		return fmt.Errorf("UpdateUserRole: failed to retrieve rows affected: %w", err)
	}

	if rowsAffected == 0 {

		return &_AdminException.UnChangeRole{
			UserID: userID,
			Role:   role,
		}
	}

	return nil
}
