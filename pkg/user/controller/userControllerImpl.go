package controller

import (
	_userService "github.com/onizukazaza/tar-ecom-api/pkg/user/service"
)

type userControllerImpl struct {
	userService _userService.UserService
}

func NewUserControllerImpl( 
	userService	_userService.UserService ,
) UserController {
	return &userControllerImpl{userService}
}