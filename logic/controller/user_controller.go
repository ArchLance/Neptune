package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"neptune/logic/data/request"
	"neptune/logic/service"
	myerrors "neptune/utils/errors"
	"neptune/utils/rsp"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}

func (controller *UserController) Update(ctx *gin.Context) {
	log.Info("controller: 创建管理员")
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBind(&updateUserRequest)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("controller: 获取更新用户参数失败 -> %w", err)})
		return
	}

	err = controller.UserService.Update(updateUserRequest)
	if err != nil {
		rsp.ErrRsp(ctx, err)
		return
	}
	rsp.SuccessRspWithNoData(ctx)
}

func (controller *UserController) Login(ctx *gin.Context) {
	log.Info("controller: 登录")
	loginRequest := request.UserLoginRequest{}
	err := ctx.ShouldBind(&loginRequest)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: err})
		return
	}
	user, err := controller.UserService.Login(loginRequest)
	if err != nil {
		var tokenInvalid myerrors.TokenInvalidErr
		if errors.As(err, &tokenInvalid) {
			rsp.ErrRsp(ctx, myerrors.TokenInvalidErr{Err: err})
			return
		}
		rsp.ErrRsp(ctx, myerrors.LoginFailed{Err: err})
		return
	}
	rsp.SuccessRsp(ctx, user)
}
