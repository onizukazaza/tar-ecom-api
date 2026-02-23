package model

type (
	SetRoleReq struct {
		ID   string `json:"id" validate:"required,uuid"` 
		Role string `json:"role" validate:"required,oneof=admin seller buyer"` 
	}
)
