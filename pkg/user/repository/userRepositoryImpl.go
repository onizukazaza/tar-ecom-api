package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/gofiber/fiber/v2/log"
	"github.com/onizukazaza/tar-ecom-api/entities"
	_userlistexception "github.com/onizukazaza/tar-ecom-api/pkg/user/exception"
	"fmt"
)

type userRepositoryImpl struct {
	db     *sqlx.DB
}


func NewUserRepositoryImpl(db *sqlx.DB ) *userRepositoryImpl {
	return &userRepositoryImpl{db: db}
}


func (r *userRepositoryImpl) Listing() ([]*entities.User, error) {
	userList := make([]*entities.User, 0)

	query := "SELECT * FROM users"

	err := r.db.Select(&userList, query)
	if err != nil {
		log.Errorf("Failed to list users: %v", err)
		return nil, &_userlistexception.UserListing{}
	}

	return userList, nil
}

func (r *userRepositoryImpl) CreateUser(user *entities.User) error {
	query := `
        INSERT INTO users (id, username, password, email, role, profile_image, created_at, updated_at)
        VALUES (:id, :username,  :password, :email, :role, :profile_image, :created_at, :updated_at)
    `
	_, err := r.db.NamedExec(query, user)
	if err != nil {
		log.Errorf("Failed to create user: %v", err)
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *userRepositoryImpl) FindUserByID(id string) (*entities.User, error) {
	query := "SELECT * FROM users WHERE id = $1"

	var user entities.User
	err := r.db.Get(&user, query, id)
	if err != nil {
		log.Errorf("Failed to find user by ID: %v", err)
		return nil, &_userlistexception.UserNotFound{}
	}
	return &user, nil
}

func (r *userRepositoryImpl) EditUser(user *entities.User) error {
	query := `
        UPDATE users 
        SET username = :username, 
            email = :email, 
            role = :role, 
            profile_image = :profile_image, 
            updated_at = :updated_at 
        WHERE id = :id
    `
	_, err := r.db.NamedExec(query, user)
	if err != nil {
		log.Errorf("Failed to update user: %v", err)
		return fmt.Errorf("failed to update user: %w", err)
	}
	return nil
}
