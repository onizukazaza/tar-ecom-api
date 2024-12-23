package entities

import (
	"github.com/google/uuid"
	"time"
)

type Address struct {
	ID            uuid.UUID `db:"id"`
	UserID        uuid.UUID `db:"user_id"`
	RecipientName string    `db:"recipient_name"`
	Province      string    `db:"province"`
	District      string    `db:"district"`
	SubDistrict   string    `db:"subdistrict"`
	Postal        string    `db:"postal"`
	AddressLine   string    `db:"address_line"`
	Contact       string    `db:"contact"`
	Favorite      bool      `db:"favorite"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`

	User []User `db:"-"`
}
