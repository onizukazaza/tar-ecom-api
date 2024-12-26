package service

import (
	_adminRepository "github.com/onizukazaza/tar-ecom-api/pkg/admin/repository"
	_userModel "github.com/onizukazaza/tar-ecom-api/pkg/user/model"
)

type adminServiceImpl struct {
	adminRepository _adminRepository.AdminRepository
}


func NewAdminServiceImpl(adminRepository _adminRepository.AdminRepository) AdminService {
	return &adminServiceImpl{adminRepository: adminRepository}
}


func (s *adminServiceImpl) Listing() ([]*_userModel.User, error) {

	userList, err := s.adminRepository.Listing()
	if err != nil {
		return nil, err
	}

	adminModelList := make([]*_userModel.User, 0)
	for _, user := range userList {
		adminModelList = append(adminModelList, user.ToUserModel())
	}

	return adminModelList, nil
}
