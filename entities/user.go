package entities

import (
	"time"

	"github.com/google/uuid"
	_adminModel "github.com/onizukazaza/tar-ecom-api/pkg/admin/model"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Username     string    `db:"username"`
	Lastname     string    `db:"lastname"`
	Email        string    `db:"email"`
	Password     string    `db:"password"`
	Role         string    `db:"role"`
	ProfileImage string    `db:"profile_image"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

func (i *User) ToItemModel() *_adminModel.User {
	return &_adminModel.User{
		ID:           i.ID.String(),
		Username:     i.Username,
		Lastname:     i.Lastname,
		Password:     i.Password,
		Email:        i.Email,
		Role:         i.Role,
		ProfileImage: i.ProfileImage,
	}
}
