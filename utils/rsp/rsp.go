package rsp

import (
	"errors"
	errorType "neptune/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func errRsp(c *gin.Context, codeErr int, codeErrMsg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": codeErr,
		"msg":  codeErrMsg,
	})
}

func ErrRsp(c *gin.Context, err error) {
	//var requestErr errorType.RequestErr
	var tokenInvalidErr errorType.TokenInvalidErr
	var loginFailed errorType.LoginFailed
	var paramErr errorType.ParamErr
	var dbErr errorType.DbErr
	var logicError errorType.LogicErr
	var notFoundErr errorType.NotFoundErr
	var existErr errorType.ExistErr
	var uploadErr errorType.UploadError
	var requestErr errorType.RequestErr
	var permissionDeniedErr errorType.PermissionDeniedError
	switch {
	case errors.As(err, &tokenInvalidErr):
		errRsp(c, errorType.CodeErrTokenInvalid, err.Error())
	case errors.As(err, &loginFailed):
		errRsp(c, errorType.CodeLoginFailed, err.Error())
	case errors.As(err, &paramErr):
		errRsp(c, errorType.CodeParamInvalid, err.Error())
	case errors.As(err, &notFoundErr):
		errRsp(c, errorType.CodeDataNotFound, err.Error())
	case errors.As(err, &existErr):
		errRsp(c, errorType.CodeDataExist, err.Error())
	case errors.As(err, &dbErr):
		errRsp(c, errorType.CodeDbError, err.Error())
	case errors.As(err, &logicError):
		errRsp(c, errorType.CodeLogicError, err.Error())
	case errors.As(err, &uploadErr):
		errRsp(c, errorType.CodeUploadError, err.Error())
	case errors.As(err, &requestErr):
		errRsp(c, errorType.CodeRequestError, err.Error())
	case errors.As(err, &permissionDeniedErr):
		errRsp(c, errorType.CodePermissionDeny, err.Error())
	default:
		errRsp(c, errorType.CodeUnknown, err.Error())
	}
}

func SuccessRsp(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	})
}

func SuccessRspWithNoData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}
