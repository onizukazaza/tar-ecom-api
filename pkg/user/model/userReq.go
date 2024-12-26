package model

type CreateUserReq struct {
	Username     string `json:"username" validate:"required,max=50"` 
	Password     string `json:"password" validate:"required,min=8,max=20"` 
	Email        string `json:"email" validate:"required,email,max=50"` 
	Role         string `json:"role" validate:"required,oneof=admin seller buyer"` 
	ProfileImage string `json:"profile_image" validate:"max=255"` 
}

type EditUserReq struct {
	ID           string `json:"id" validate:"required"` 
	Username     string `json:"username" validate:"max=50"` 
	Email        string `json:"email" validate:"omitempty,email,max=100"` 
	Role         string `json:"role" validate:"omitempty,oneof=admin seller buyer"` 
	ProfileImage string `json:"profile_image" validate:"max=255"`
}
