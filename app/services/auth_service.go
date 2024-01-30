package services

import (
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/http/requests"
	"goravel/app/models"
)

type AuthService struct {
	// Service-specific dependencies can be added here
}

type UserResponse struct {
	Uuid       string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Authent    string `json:"authent"`
	Avatar     string `json:"avatar"`
	AvatarPath string `json:"avatar_path"`
}

func NewAuthService() *AuthService {
	return &AuthService{
		// Initialize any dependencies if needed
	}
}

func (a *AuthService) LoginProcess(payload requests.UserLoginRequest) (bool, models.User, UserResponse) {
	var userModel models.User

	err := facades.Orm().Query().Where("email", payload.Email).Where("status", 1).First(&userModel)

	if err != nil {
		facades.Log().Error(err)
		return false, models.User{}, UserResponse{}
	}

	check, err := helpers.CheckPasswordHash(payload.Password, userModel.Password)

	if err != nil {
		return false, models.User{}, UserResponse{}
	}

	if !check {
		return false, models.User{}, UserResponse{}
	}

	result := UserResponse{
		Uuid:    userModel.Uuid,
		Name:    userModel.Name,
		Email:   userModel.Email,
		Authent: userModel.Authent,
		Avatar:  userModel.Avatar,
	}

	return true, userModel, result
}
