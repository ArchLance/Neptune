package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"neptune/global"
	"neptune/logic/data/request"
	"neptune/logic/service"
	myerrors "neptune/utils/errors"
	"neptune/utils/file"
	img "neptune/utils/image"
	"neptune/utils/rsp"
	"os"
	"path"
	"strconv"
	"strings"
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

	err = controller.UserService.Update(&updateUserRequest)
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
	user, err := controller.UserService.Login(&loginRequest)
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

func (controller *UserController) UploadAvatar(ctx *gin.Context) {

	f, err := ctx.FormFile("imgfile")
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.UploadError{Err: fmt.Errorf("上传失败!")})
		return
	} else {
		fileExt := strings.ToLower(path.Ext(f.Filename))
		if !img.CheckImg(f) {
			rsp.ErrRsp(ctx, myerrors.UploadError{Err: fmt.Errorf("上传失败!只允许png,jpg,jpeg文件")})
			return
		}
		// 获得userId
		userId := ctx.PostForm("userId")
		fileName := userId
		id, _ := strconv.Atoi(userId)
		userResponse, err := controller.UserService.GetById(id)
		if err != nil {
			rsp.ErrRsp(ctx, myerrors.UploadError{Err: fmt.Errorf("获取用户信息失败")})
			return
		}
		userRequest := request.UpdateUserRequest{
			UserId:   id,
			Avatar:   fmt.Sprintf("%s%s", fileName, fileExt),
			UserName: userResponse.UserName,
			Account:  userResponse.Account,
			Email:    userResponse.Email,
			Role:     userResponse.Role,
		}
		err = controller.UserService.Update(&userRequest)
		if err != nil {
			rsp.ErrRsp(ctx, myerrors.UploadError{Err: fmt.Errorf("数据库更新用户信息失败")})
			return
		}
		fileDir := global.ServerConfig.BaseConf.Upload.Avatar
		isExist, _ := file.IsFileExist(fileDir)
		if !isExist {
			err := os.Mkdir(fileDir, os.ModePerm)
			if err != nil {
				rsp.ErrRsp(ctx, myerrors.UploadError{Err: fmt.Errorf("创建文件夹失败")})
				return
			}
		}
		filepath := fmt.Sprintf("%s%s%s", fileDir, fileName, fileExt)
		err = ctx.SaveUploadedFile(f, filepath)
		if err != nil {
			rsp.ErrRsp(ctx, myerrors.UploadError{Err: fmt.Errorf("服务器保存图片失败")})
			return
		}
		rsp.SuccessRsp(ctx, gin.H{
			"path": fmt.Sprintf("%s%s", fileName, fileExt),
		})
	}
}
