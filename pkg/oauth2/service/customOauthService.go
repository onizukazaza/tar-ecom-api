package service

import (
	"github.com/onizukazaza/tar-ecom-api/pkg/oauth2/model"
)

type OAuth2Service interface {
	Login(email, password string) (*model.LoginResponse, error)
	Logout(token string) error
}
