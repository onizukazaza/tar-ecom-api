package model

import "time"

//Req  represents
type CreateAddressReq struct {
	UserID        string    `json:"user_id" validate:"required,max=64"`  
	RecipientName string    `json:"recipient_name" validate:"required,max=126"` 
	Province      string    `json:"province" validate:"required,max=64"` 
	District      string    `json:"district" validate:"required,max=64"` 
	SubDistrict   string    `json:"subdistrict" validate:"required,max=64"` 
	Postal        string    `json:"postal" validate:"required,max=5,numeric"`
	AddressLine   string    `json:"address_line" validate:"required,max=255"` 
	Contact       string    `json:"contact" validate:"required,max=10,numeric"` 
	Favorite      bool      `json:"favorite"`
	IsDefault     bool      `json:"is_default"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}



type EditAddressReq struct {
    ID            string `json:"id" validate:"required"`            
    UserID        string `json:"-"`                             
    RecipientName string `json:"recipient_name,omitempty" validate:"max=126"` 
    Province      string `json:"province,omitempty" validate:"max=64"` 
    District      string `json:"district,omitempty" validate:"max=64"`  
    SubDistrict   string `json:"subdistrict,omitempty" validate:"max=64"`
    Postal        string `json:"postal,omitempty" validate:"max=5,numeric"`  
    AddressLine   string `json:"address_line,omitempty" validate:"max=255"`  
    Contact       string `json:"contact,omitempty" validate:"max=10,numeric"`
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
