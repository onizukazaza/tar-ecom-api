package entities

import (
	"time"

	"github.com/google/uuid"
	_userModel "github.com/onizukazaza/tar-ecom-api/pkg/user/model"
	
)

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleSeller Role = "seller"
	RoleBuyer  Role = "buyer"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Username     string    `db:"username"`
	// Lastname     string    `db:"lastname"`
	Email        string    `db:"email"`
	Password     string    `db:"password"`
	Role         Role      `db:"role"`
	ProfileImage string    `db:"profile_image"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (u *User) ToUserModel() *_userModel.User {
    return &_userModel.User{
        ID:           u.ID.String(),
        Username:     u.Username,
        // Lastname:     u.Lastname,
        Password:     u.Password, 
        Email:        u.Email,
        Role:         string(u.Role),
        ProfileImage: u.ProfileImage,
    }
}
