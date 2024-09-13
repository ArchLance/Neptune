package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"neptune/global"
	"neptune/logic/data/request"
	"neptune/logic/service"
	email "neptune/utils/email"
	myerrors "neptune/utils/errors"
	"neptune/utils/file"
	"neptune/utils/hash"
	img "neptune/utils/image"
	"neptune/utils/random"
	"neptune/utils/rsp"
	jwt "neptune/utils/token"
	"os"
	"path"
	"strings"
	"time"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}

func (c *UserController) Update(ctx *gin.Context) {
	log.Info("controller: 更新用户")
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBind(&updateUserRequest)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("controller: 获取更新用户参数失败 -> %w", err)})
		return
	}
	// 防止构造随便一个用户的id就可以修改其密码
	claims := jwt.GetClaims(ctx)
	updateUserRequest.UserId = claims.UserID
	err = c.UserService.Update(&updateUserRequest)
	if err != nil {
		rsp.ErrRsp(ctx, err)
		return
	}
	rsp.SuccessRspWithNoData(ctx)
}

func (c *UserController) Login(ctx *gin.Context) {
	log.Info("controller: 登录")
	loginRequest := request.UserLoginRequest{}
	err := ctx.ShouldBind(&loginRequest)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: err})
		return
	}
	user, err := c.UserService.Login(&loginRequest)
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

func (c *UserController) UploadAvatar(ctx *gin.Context) {

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
		fileName := hash.Md5str(fmt.Sprintf("%s%s", f.Filename, time.Now().String()))
		// 防止构造随便一个用户的id就可以修改其密码
		claims := jwt.GetClaims(ctx)
		if claims == nil {
			rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("获取用户信息失败")})
			return
		}
		userResponse, err := c.UserService.GetById(claims.UserID)
		if err != nil {
			rsp.ErrRsp(ctx, myerrors.UploadError{Err: fmt.Errorf("获取用户信息失败")})
			return
		}
		userRequest := request.UpdateUserRequest{
			UserId:   claims.UserID,
			Avatar:   fmt.Sprintf("%s%s", fileName, fileExt),
			UserName: userResponse.UserName,
			Account:  userResponse.Account,
			Email:    userResponse.Email,
			Role:     userResponse.Role,
		}
		err = c.UserService.Update(&userRequest)
		if err != nil {
			rsp.ErrRsp(ctx, myerrors.UploadError{Err: fmt.Errorf("数据库更新用户信息失败")})
			return
		}
		fileDir := global.ServerConfig.BaseConfig.Upload.Avatar
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

func (c *UserController) ChangePassword(ctx *gin.Context) {
	log.Info("用户修改密码")
	changePassword := request.UserChangePasswordRequest{}
	err := ctx.ShouldBindJSON(&changePassword)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("参数错误 -> %w", err)})
		return
	}
	claims := jwt.GetClaims(ctx)
	changePassword.UserId = claims.UserID
	err = c.UserService.ChangePassword(&changePassword)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: err})
		return
	}
	rsp.SuccessRspWithNoData(ctx)
}

func (c *UserController) GenerateCode(ctx *gin.Context) {
	log.Info("生成验证码")
	// 查询用户
	claims := jwt.GetClaims(ctx)
	if claims == nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("获取用户信息失败")})
		return
	}
	userResponse, err := c.UserService.GetById(claims.UserID)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("获取用户信息失败")})
		return
	}
	// TODO： 获取redis中的验证码的过期时间，避免短时间内重复生成：
	ok := global.Redis.Get(ctx, userResponse.Email)
	if ok != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("请不要频繁发送验证码")})
		return
	}

	code := random.GenValidateCode(6)
	err = global.Redis.Set(ctx, userResponse.Email, code, 180*time.Second).Err()
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.DbErr{Err: err})
		return
	}
	emailContent := fmt.Sprintf("您的验证码是：%s，此验证码3分钟有效。", code)
	err = email.SendEmail(userResponse.Email, "邮箱绑定验证码", emailContent)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("发送邮件失败: %w", err)})
		return
	}
	rsp.SuccessRspWithNoData(ctx)
}

func (c *UserController) CheckCode(ctx *gin.Context) {
	log.Info("校验验证码")
	requestCode := ctx.Query("code")
	if requestCode == "" {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("验证码不能为空")})
		return
	}
	claims := jwt.GetClaims(ctx)
	if claims == nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("获取用户信息失败")})
		return
	}
	userResponse, err := c.UserService.GetById(claims.UserID)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("获取用户信息失败")})
		return
	}

	code, err := global.Redis.Get(ctx, userResponse.Email).Result()
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.NotFoundErr{Err: fmt.Errorf("验证码错误")})
		return
	}
	if code != requestCode {
		rsp.ErrRsp(ctx, myerrors.NotFoundErr{Err: fmt.Errorf("验证码错误")})
		return
	}
	rsp.SuccessRspWithNoData(ctx)
}
