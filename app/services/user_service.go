package services

import (
	"github.com/google/uuid"
	"github.com/goravel/framework/facades"
	"goravel/app/helpers"
	"goravel/app/http/requests"
	"goravel/app/models"
	"strconv"
)

type UserService struct {
	// Service-specific dependencies can be added here
}

func NewUserService() *UserService {
	return &UserService{
		// Initialize any dependencies if needed
	}
}

/*
*
Fetch Data User
*/

func (s *UserService) Fetch() []models.User {
	var user []models.User

	err := facades.Orm().Query().Order("id desc").Find(&user)

	if err != nil {
		facades.Log().Error(err)
	}

	return user
}

/*
*
Get Data By UUID
*/

func (s *UserService) GetByUuid(uuid string) models.User {
	var user models.User

	err := facades.Orm().Query().Where("uuid", uuid).First(&user)

	if err != nil {
		facades.Log().Error(err)
	}

	return user
}

/**
Get Data By Id
*/

func (s *UserService) GetById(id string) models.User {
	var user models.User

	err := facades.Orm().Query().FindOrFail(&user, id)

	if err != nil {
		facades.Log().Error(err)
	}

	return user
}

/**
Store New Data
*/

func (s *UserService) Store(data requests.UserStoreRequest) {
	var user models.User

	hashedPassword, err := helpers.HashPassword(data.Password)

	if err != nil {
		facades.Log().Error(err)
	}

	user.Uuid = uuid.NewString()
	user.Name = data.Name
	user.Email = data.Email
	user.Password = hashedPassword
	user.Authent = data.Authent
	user.Status, _ = strconv.Atoi(data.Status)

	err = facades.Orm().Query().Save(&user)

	if err != nil {
		facades.Log().Error(err)
	}
}

/**
Update Data
*/

func (s *UserService) Update(data requests.UserUpdateRequest, uuid string) {
	var user models.User

	err := facades.Orm().Query().Where("uuid", uuid).First(&user)

	if err != nil {
		facades.Log().Error(err)
	}

	user.Name = data.Name
	user.Email = data.Email
	user.Authent = data.Authent
	user.Status, _ = strconv.Atoi(data.Status)

	err = facades.Orm().Query().Save(&user)

	if err != nil {
		facades.Log().Error(err)
	}
}

func (s *UserService) destroy(uuid string) {
	_, err := facades.Orm().Query().Where("uuid", uuid).Delete(&models.User{})

	if err != nil {
		facades.Log().Error(err)
	}
}

func (s *UserService) CheckEmail(email string) bool {
	var model models.User
	var count int64

	err := facades.Orm().Query().Model(model).Where("email", email).Count(&count)

	if err != nil {
		facades.Log().Error(err)
	}

	if count < 1 {
		return true
	}

	return false
}

func (s *UserService) Register(payload requests.UserRegisterRequest) {
	var model models.User

	hashedPassword, err := helpers.HashPassword(payload.Password)

	if err != nil {
		facades.Log().Error(err)
	}

	model.Uuid = uuid.NewString()
	model.Name = payload.Name
	model.Email = payload.Email
	model.Password = hashedPassword
	model.Authent = "user"
	model.Status = 1

	err = facades.Orm().Query().Save(&model)

	if err != nil {
		facades.Log().Error(err)
	}
}
