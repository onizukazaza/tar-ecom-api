package entities

import (
	"github.com/google/uuid"
	"time"
	_addressModel "github.com/onizukazaza/tar-ecom-api/pkg/address/model"
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


func (a *Address) ToModel() *_addressModel.Address {
	return &_addressModel.Address{
		ID:            a.ID.String(),
		UserID:        a.UserID.String(),
		RecipientName: a.RecipientName,
		Province:      a.Province,
		District:      a.District,
		SubDistrict:   a.SubDistrict,
		Postal:        a.Postal,
		AddressLine:   a.AddressLine,
		Contact:       a.Contact,
		Favorite:      a.Favorite,

	}
}

