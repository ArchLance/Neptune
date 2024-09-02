package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"neptune/logic/data/request"
	"neptune/logic/data/response"
	"neptune/logic/repository"
	myerrors "neptune/utils/errors"
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

func (r *UserServiceImpl) Update(user request.UpdateUserRequest) error {
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
	err = r.UserRepository.Update(&userData)
	if err != nil {
		return myerrors.DbErr{Err: err}
	}
	return nil
}

func (r *UserServiceImpl) Login(user request.UserLoginRequest) (response.UserResponse, error) {
	err := r.Validate.Struct(user)
	if err != nil {
		return response.UserResponse{}, myerrors.LoginFailed{Err: fmt.Errorf("service: 登录用户参数校验失败 -> %w", err)}
	}
	userData, err := r.UserRepository.GetByAccount(user.Account)
	if err != nil {
		return response.UserResponse{}, myerrors.LoginFailed{Err: fmt.Errorf("用户名密码错误")}
	}
	if userData.Password != user.Password {
		return response.UserResponse{}, myerrors.LoginFailed{Err: fmt.Errorf("用户名密码错误")}
	}
	token, err := jwt.GenerateToken(userData.Id, userData.Role)
	if err != nil {
		return response.UserResponse{}, myerrors.TokenInvalidErr{Err: err}
	}
	responseUser := response.UserResponse{
		UserId:   userData.Id,
		Token:    token,
		UserName: userData.UserName,
		Account:  userData.Account,
		Email:    userData.Email,
		Role:     userData.Role,
	}
	return responseUser, nil
}
