package service

import (
	_userRepository "github.com/onizukazaza/tar-ecom-api/pkg/user/repository"
	_userModel "github.com/onizukazaza/tar-ecom-api/pkg/user/model"
	"golang.org/x/crypto/bcrypt"
	"github.com/onizukazaza/tar-ecom-api/entities"
	"github.com/google/uuid"
	"fmt"
	"time"
)

type userServiceImpl struct {
	userRepository _userRepository.UserRepository
}


func NewUserServiceImpl(userRepository _userRepository.UserRepository) UserService {
	return &userServiceImpl{userRepository: userRepository}
}


func (s *userServiceImpl) Listing() ([]*_userModel.User, error) {

	userList, err := s.userRepository.Listing()
	if err != nil {
		return nil, err
	}

	adminModelList := make([]*_userModel.User, 0)
	for _, user := range userList {
		adminModelList = append(adminModelList, user.ToUserModel())
	}

	return adminModelList, nil
}

func (s *userServiceImpl) CreateUser(req *_userModel.CreateUserReq) (*_userModel.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &entities.User{
		ID:           uuid.New(),
		Username:     req.Username,
		// Lastname:     req.Lastname,
		Password:     string(hashedPassword),
		Email:        req.Email,
		Role:         entities.Role(req.Role),
		ProfileImage: req.ProfileImage,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.userRepository.CreateUser(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user.ToUserModel(), nil
}

func (s *userServiceImpl) FindUserByID(id string) (*_userModel.User, error) {
	user, err := s.userRepository.FindUserByID(id)
	if err != nil {
		return nil, err
	}
	return user.ToUserModel(), nil
}

func (s *userServiceImpl) EditUser(req *_userModel.EditUserReq) error {
	user, err := s.userRepository.FindUserByID(req.ID)
	if err != nil {
		return err
	}

	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Role != "" {
		user.Role = entities.Role(req.Role)
	}
	if req.ProfileImage != "" {
		user.ProfileImage = req.ProfileImage
	}
	user.UpdatedAt = time.Now()

	if err := s.userRepository.EditUser(user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
