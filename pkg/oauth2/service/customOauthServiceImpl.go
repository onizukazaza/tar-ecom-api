package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	// "github.com/onizukazaza/tar-ecom-api/entities"
	"github.com/onizukazaza/tar-ecom-api/pkg/oauth2/model"
	_userRepository "github.com/onizukazaza/tar-ecom-api/pkg/user/repository"
	"fmt"
)

type oauth2ServiceImpl struct {
	userRepository _userRepository.UserRepository
	secretKey      string
}

func NewOAuth2Service(userRepository _userRepository.UserRepository, secretKey string) OAuth2Service {
	return &oauth2ServiceImpl{userRepository: userRepository, secretKey: secretKey}
}

func (s *oauth2ServiceImpl) Login(email, password string) (*model.LoginResponse, error) {
	user, err := s.userRepository.FindUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	claims := jwt.MapClaims{
		"id":   user.ID.String(),
		"role": string(user.Role),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return nil, fmt.Errorf("failed to generate token")
	}

	return &model.LoginResponse{
		UserID: user.ID.String(),
		Role:   string(user.Role),
		Token:  tokenString,
	}, nil
}


func (s *oauth2ServiceImpl) Logout(token string) error {

    if token == "" {
        return fmt.Errorf("invalid token")
    }

    return nil
}
