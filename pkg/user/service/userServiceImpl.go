package service

import (
	_userRepository "github.com/onizukazaza/tar-ecom-api/pkg/user/repository"
 	   )

type userServiceImpl struct {
	userRepository _userRepository.UserRepository
}

func NewUserServiceImpl(
	userRepository _userRepository.UserRepository ,
) UserService {
	return &userServiceImpl{userRepository}
}

