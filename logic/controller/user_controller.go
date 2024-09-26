package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"neptune/global"
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
	UserService *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{
		UserService: service,
	}
}

func (c *UserController) Update(ctx *gin.Context) {
	log.Info("controller: 更新用户")
	updateUserRequest := service.UpdateUserRequest{}
	err := ctx.ShouldBind(&updateUserRequest)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: err})
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
	loginRequest := service.UserLoginRequest{}
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
		userRequest := service.UpdateUserRequest{
			UserId:   claims.UserID,
			Avatar:   fmt.Sprintf("%s%s", fileName, fileExt),
			UserName: userResponse.UserName,
			Account:  userResponse.Account,
			Email:    userResponse.Email,
			Role:     userResponse.Role,
		}
		err = c.UserService.Update(&userRequest)
		if err != nil {
			rsp.ErrRsp(ctx, myerrors.UploadError{Err: err})
			return
		}
		fileDir := global.ServerConfig.BaseConfig.Upload.Avatar
		isExist, _ := file.IsFileExist(fileDir)
		if !isExist {
			err := os.Mkdir(fileDir, os.ModePerm)
			if err != nil {
				rsp.ErrRsp(ctx, myerrors.UploadError{Err: err})
				return
			}
		}
		filepath := fmt.Sprintf("%s%s%s", fileDir, fileName, fileExt)
		err = ctx.SaveUploadedFile(f, filepath)
		if err != nil {
			rsp.ErrRsp(ctx, myerrors.UploadError{Err: err})
			return
		}
		rsp.SuccessRsp(ctx, gin.H{
			"path": fmt.Sprintf("%s%s", fileName, fileExt),
		})
	}
}

func (c *UserController) ChangePassword(ctx *gin.Context) {
	log.Info("用户修改密码")
	changePassword := service.UserChangePasswordRequest{}
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

type ExpireCode struct {
	TimeStamp int64
	Code      string
}

func (c *UserController) GenerateCode(ctx *gin.Context) {
	log.Info("生成验证码")
	// 查询用户
	requestEmail := ctx.Query("email")
	generateType := ctx.Query("type")
	if generateType == "0" {
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
		requestEmail = userResponse.Email
	}

	// TODO： 获取redis中的验证码的过期时间，避免短时间内重复生成：
	res := global.Redis.Get(ctx, requestEmail)
	if res.Val() != "" {
		var expireCode ExpireCode
		val, err := res.Bytes()
		if err != nil {
			rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("验证码信息解析失败")})
			return
		}
		err = json.Unmarshal(val, &expireCode)
		if err != nil {
			rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("验证码信息解析失败")})
			return
		}
		// 限制请求60秒才能更新一次，否则不更新
		generateTime := time.Now().Unix() - expireCode.TimeStamp
		if generateTime < 60 {
			rsp.ErrRsp(ctx, myerrors.LogicErr{Err: fmt.Errorf("请求验证码过于频繁")})
			return
		}
	}

	code := random.GenValidateCode(6)
	expireCode := ExpireCode{
		TimeStamp: time.Now().Unix(),
		Code:      code,
	}
	data, _ := json.Marshal(expireCode)
	err := global.Redis.Set(ctx, requestEmail, data, 180*time.Second).Err()
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.DbErr{Err: err})
		return
	}
	emailContent := fmt.Sprintf("您的验证码是：%s，此验证码3分钟有效。", code)
	// 起协程发邮件
	go func() {
		err := email.SendEmail(requestEmail, "邮箱绑定验证码", emailContent)
		if err != nil {
			log.Error("发送邮件失败")
		}
	}()
	rsp.SuccessRspWithNoData(ctx)
}

func checkCode(ctx *gin.Context, email string, code string) bool {
	res := global.Redis.Get(ctx, email)
	if res.Val() == "" {
		rsp.ErrRsp(ctx, myerrors.NotFoundErr{Err: fmt.Errorf("验证码错误")})
		return false
	}
	var expireCode ExpireCode
	val, err := res.Bytes()
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("验证码信息解析失败")})
		return false
	}
	err = json.Unmarshal(val, &expireCode)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("验证码信息解析失败")})
		return false
	}
	if expireCode.Code != code {
		rsp.ErrRsp(ctx, myerrors.NotFoundErr{Err: fmt.Errorf("验证码错误")})
		return false
	}
	return true
}

func (c *UserController) CheckCode(ctx *gin.Context) {
	log.Info("校验验证码")
	requestCode := ctx.Query("code")
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

	if checkCode(ctx, userResponse.Email, requestCode) {
		rsp.SuccessRspWithNoData(ctx)
	}

}

func (c *UserController) UpdateEmail(ctx *gin.Context) {
	log.Info("更新邮箱")
	changeEmail := service.UserChangeEmailRequest{}
	err := ctx.ShouldBindJSON(&changeEmail)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("参数错误 -> %w", err)})
		return
	}
	claims := jwt.GetClaims(ctx)
	changeEmail.UserId = claims.UserID
	// 如果验证码校验不通过
	if !checkCode(ctx, changeEmail.Email, changeEmail.Code) {
		return
	}
	err = c.UserService.ChangeEmail(&changeEmail)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: err})
		return
	}
	rsp.SuccessRspWithNoData(ctx)
}
