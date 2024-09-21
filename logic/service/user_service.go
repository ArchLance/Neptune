package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"neptune/logic/repository"
	myerrors "neptune/utils/errors"
	"neptune/utils/hash"
	jwt "neptune/utils/token"
)

type UserService struct {
	UserRepository *repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(repository *repository.UserRepository, validate *validator.Validate) *UserService {
	return &UserService{
		UserRepository: repository,
		Validate:       validate,
	}
}

type UserLoginRequest struct {
	Account  string `validate:"required,max=64,min=1"  json:"account"`
	Password string `validate:"required,max=64,min=1" json:"password"`
}

type UpdateUserRequest struct {
	UserId   uint   `json:"userid"`
	Avatar   string `json:"avatar"`
	UserName string `validate:"required,max=64,min=1" json:"username"`
	Account  string `validate:"required,max=64,min=1"  json:"account"`
	Email    string `validate:"required,max=64,min=1" json:"email"`
	Role     string `validate:"required,max=20,min=1" json:"role"`
}

type UserChangePasswordRequest struct {
	UserId      uint   `validate:"max=64,min=1"  json:"user_id"`
	OldPassword string `validate:"required,max=64,min=1" json:"old_password"`
	NewPassword string `validate:"required,max=64,min=1" json:"new_password"`
}

type UserChangeEmailRequest struct {
	UserId uint   `validate:"max=64,min=1"  json:"user_id"`
	Email  string `validate:"required,max=64,min=1" json:"email"`
	Code   string `validate:"required,max=64,min=1" json:"code"`
}
type UserLoginResponse struct {
	UserId   int    `json:"userid"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Account  string `json:"account"`
	Role     string `json:"role"`
}

type UserResponse struct {
	UserId   int    `json:"userid"`
	Avatar   string `json:"avatar"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Account  string `json:"account"`
	Role     string `json:"role"`
}

func (r *UserService) GetById(id uint) (UserResponse, error) {
	userData, err := r.UserRepository.GetById(id)
	if err != nil {
		return UserResponse{}, myerrors.NotFoundErr{Err: err}
	}
	userResponse := UserResponse{
		UserId:   userData.Id,
		UserName: userData.UserName,
		Account:  userData.Account,
		Avatar:   userData.Avatar,
		Role:     userData.Role,
		Email:    userData.Email,
	}
	return userResponse, nil
}

func (r *UserService) Update(user *UpdateUserRequest) error {
	err := r.Validate.Struct(user)
	if err != nil {
		return myerrors.ParamErr{Err: fmt.Errorf("service: 更新用户参数校验失败 -> %w", err)}
	}
	userData, err := r.UserRepository.GetById(user.UserId)
	if err != nil {
		return myerrors.NotFoundErr{Err: err}
	}
	userData.UserName = user.UserName
	userData.Account = user.Account
	userData.Email = user.Email
	userData.Avatar = user.Avatar
	userData.Role = user.Role
	err = r.UserRepository.Update(&userData)
	if err != nil {
		return myerrors.DbErr{Err: err}
	}
	return nil
}

func (r *UserService) Login(user *UserLoginRequest) (UserLoginResponse, error) {
	err := r.Validate.Struct(user)
	if err != nil {
		return UserLoginResponse{}, myerrors.LoginFailed{Err: fmt.Errorf("service: 登录用户参数校验失败 -> %w", err)}
	}
	userData, err := r.UserRepository.GetByAccount(user.Account)
	if err != nil {
		return UserLoginResponse{}, myerrors.LoginFailed{Err: fmt.Errorf("用户名密码错误")}
	}
	password := hash.SHA256DoubleString(user.Password, false)
	if userData.Password != password {
		return UserLoginResponse{}, myerrors.LoginFailed{Err: fmt.Errorf("用户名密码错误")}
	}
	token, err := jwt.GenerateToken(userData.Id, userData.Role)
	if err != nil {
		return UserLoginResponse{}, myerrors.TokenInvalidErr{Err: err}
	}
	responseUser := UserLoginResponse{
		UserId:   userData.Id,
		Token:    token,
		UserName: userData.UserName,
		Account:  userData.Account,
		Email:    userData.Email,
		Role:     userData.Role,
		Avatar:   userData.Avatar,
	}
	return responseUser, nil
}

func (r *UserService) ChangePassword(user *UserChangePasswordRequest) error {
	err := r.Validate.Struct(user)
	if err != nil {
		return myerrors.ParamErr{Err: fmt.Errorf("service: 修改密码参数校验失败 -> %w", err)}
	}
	userData, err := r.UserRepository.GetById(user.UserId)
	if err != nil {
		return myerrors.PermissionDeniedError{Err: fmt.Errorf("权限校验失败 -> %w", err)}
	}
	//// 判断密码是否正确  此处要保证原密码校验通过才能修改密码
	oldPassword := hash.SHA256DoubleString(user.OldPassword, false)
	if userData.Password != oldPassword {
		return myerrors.PermissionDeniedError{Err: fmt.Errorf("原密码错误")}
	}
	newPassword := hash.SHA256DoubleString(user.NewPassword, false)
	userData.Password = newPassword
	err = r.UserRepository.Update(&userData)
	if err != nil {
		return myerrors.DbErr{Err: fmt.Errorf("修改密码失败 -> %w", err)}
	}
	return nil
}

func (r *UserService) ChangeEmail(user *UserChangeEmailRequest) error {
	err := r.Validate.Struct(user)
	if err != nil {
		return myerrors.ParamErr{Err: fmt.Errorf("service: 修改邮箱参数校验失败 -> %w", err)}
	}
	userData, err := r.UserRepository.GetById(user.UserId)
	if err != nil {
		return myerrors.PermissionDeniedError{Err: fmt.Errorf("权限校验失败 -> %w", err)}
	}
	userData.Email = user.Email
	err = r.UserRepository.Update(&userData)
	if err != nil {
		return myerrors.DbErr{Err: fmt.Errorf("修改密码失败 -> %w", err)}
	}
	return nil
}
