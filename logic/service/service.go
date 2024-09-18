package service

import (
	"neptune/logic/data/request"
	"neptune/logic/data/response"
)

type ManagerService interface {
	Create(manager request.CreateManagerRequest) error
	Update(manager request.UpdateManagerRequest) error
	Delete(id int) error
	GetById(id int) (response.ManagerResponse, error)
	GetAll() ([]response.ManagerResponse, error)
}

type UserService interface {
	GetById(id uint) (response.UserResponse, error)
	Update(user *request.UpdateUserRequest) error
	Login(user *request.UserLoginRequest) (response.UserLoginResponse, error)
	ChangePassword(user *request.UserChangePasswordRequest) error
	ChangeEmail(user *request.UserChangeEmailRequest) error
}
