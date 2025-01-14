package service

import (
	"fmt"
	_adminRepository "github.com/onizukazaza/tar-ecom-api/pkg/admin/repository"
	_userRepository "github.com/onizukazaza/tar-ecom-api/pkg/user/repository"
	_adminModel "github.com/onizukazaza/tar-ecom-api/pkg/admin/model"
	_userModel  "github.com/onizukazaza/tar-ecom-api/pkg/user/model"
)

type adminServiceImpl struct {
	adminRepository _adminRepository.AdminRepository
	userRepository  _userRepository.UserRepository
}

func NewAdminServiceImpl(adminRepository _adminRepository.AdminRepository, userRepository _userRepository.UserRepository) AdminService {
	return &adminServiceImpl{
		adminRepository: adminRepository,
		userRepository:  userRepository,
	}
}

func (s *adminServiceImpl) SetRole(req *_adminModel.SetRoleReq) error {
	// ตรวจสอบว่าผู้ใช้มีอยู่จริง
	_ , err := s.userRepository.FindUserByID(req.ID)
	if err != nil {
		return err
	}

	// if user.Role == req.Role {
	// 	return fmt.Errorf("SetRole: no changes to role")
	// }

	err = s.adminRepository.UpdateUserRole(req.ID, req.Role)
	if err != nil {
		return fmt.Errorf("SetRole: failed to update user role: %w", err)
	}

	return nil
}

func (s *adminServiceImpl) Listing() ([]*_userModel.User, error) {

	userList, err := s.userRepository.Listing()
	if err != nil {
		return nil, err
	}

	adminModelList := make([]*_userModel.User, 0)
	for _, user := range userList {
		adminModelList = append(adminModelList, user.ToModel())
	}

	return adminModelList, nil
}


