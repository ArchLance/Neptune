package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"neptune/logic/data/request"
	"neptune/logic/data/response"
	"neptune/logic/repository"
	myerrors "neptune/utils/errors"
	"neptune/utils/hash"
	jwt "neptune/utils/token"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func NewUserServiceImpl(repository repository.UserRepository, validate *validator.Validate) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: repository,
		Validate:       validate,
	}
}

func (r *UserServiceImpl) GetById(id int) (response.UserResponse, error) {
	userData, err := r.UserRepository.GetById(id)
	if err != nil {
		return response.UserResponse{}, myerrors.NotFoundErr{Err: err}
	}
	userResponse := response.UserResponse{
		UserId:   userData.Id,
		UserName: userData.UserName,
		Account:  userData.Account,
		Avatar:   userData.Avatar,
		Role:     userData.Role,
		Email:    userData.Email,
	}
	return userResponse, nil
}

func (r *UserServiceImpl) Update(user *request.UpdateUserRequest) error {
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

func (r *UserServiceImpl) Login(user *request.UserLoginRequest) (response.UserLoginResponse, error) {
	err := r.Validate.Struct(user)
	if err != nil {
		return response.UserLoginResponse{}, myerrors.LoginFailed{Err: fmt.Errorf("service: 登录用户参数校验失败 -> %w", err)}
	}
	userData, err := r.UserRepository.GetByAccount(user.Account)
	if err != nil {
		return response.UserLoginResponse{}, myerrors.LoginFailed{Err: fmt.Errorf("用户名密码错误")}
	}
	password := hash.SHA256DoubleString(user.Password, false)
	if userData.Password != password {
		return response.UserLoginResponse{}, myerrors.LoginFailed{Err: fmt.Errorf("用户名密码错误")}
	}
	token, err := jwt.GenerateToken(userData.Id, userData.Role)
	if err != nil {
		return response.UserLoginResponse{}, myerrors.TokenInvalidErr{Err: err}
	}
	responseUser := response.UserLoginResponse{
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
