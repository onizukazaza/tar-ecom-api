package model

import "time"

//Req  represents
type CreateAddressReq struct {
	UserID        string    `json:"user_id"`
	RecipientName string    `json:"recipient_name" validate:"required"`
	Province      string    `json:"province" validate:"required"`
	District      string    `json:"district" validate:"required"`
	SubDistrict   string    `json:"subdistrict" validate:"required"`
	Postal        string    `json:"postal" validate:"required"`
	AddressLine   string    `json:"address_line" validate:"required"`
	Contact       string    `json:"contact" validate:"required"`
	Favorite      bool      `json:"favorite"`
	IsDefault     bool      `json:"is_default"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type EditAddressReq struct {
    ID            string `json:"id" validate:"required"`
    UserID        string `json:"-"`
    RecipientName string `json:"recipient_name,omitempty"`
    Province      string `json:"province,omitempty"`
    District      string `json:"district,omitempty"`
    SubDistrict   string `json:"subdistrict,omitempty"`
    Postal        string `json:"postal,omitempty"`
    AddressLine   string `json:"address_line,omitempty"`
    Contact       string `json:"contact,omitempty"`
    Favorite      bool   `json:"favorite,omitempty"`
}

//Response represents
type Address struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	RecipientName string `json:"recipient_name"`
	Province      string `json:"province"`
	District      string `json:"district"`
	SubDistrict   string `json:"subdistrict"`
	Postal        string `json:"postal"`
	AddressLine   string `json:"address_line"`
	Contact       string `json:"contact"`
	Favorite      bool   `json:"favorite"`
}
