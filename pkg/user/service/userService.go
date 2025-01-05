package service

import (
    _userModel "github.com/onizukazaza/tar-ecom-api/pkg/user/model"
)
type UserService interface {
	Listing() ([] *_userModel.User, error)
	CreateUser(req *_userModel.CreateUserReq) (*_userModel.User, error)
	FindUserByID(id string) (*_userModel.User, error)
	EditUser(req *_userModel.EditUserReq) error 
	IsUserExists(email string) (bool, error)
}