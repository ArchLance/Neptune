package rsp

import (
	"errors"
	"net/http"
	errorType "student_manage/utils/errors"

	"github.com/gin-gonic/gin"
)

func errRsp(c *gin.Context, codeErr int, codeErrMsg string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"code": codeErr,
		"msg":  codeErrMsg,
	})
}

func ErrRsp(c *gin.Context, err error) {
	var requestErr errorType.RequestErr
	var paramErr errorType.ParamErr
	var dbErr errorType.DbErr
	var logicError errorType.LogicErr
	var notFoundErr errorType.NotFoundErr
	var existErr errorType.ExistErr
	switch {
	case errors.As(err, &requestErr):
		errRsp(c, errorType.CodeRequestFailed, "请求错误")
	case errors.As(err, &paramErr):
		errRsp(c, errorType.CodeParamInvalid, "参数错误")
	case errors.As(err, &notFoundErr):
		errRsp(c, errorType.CodeDataNotFound, "数据查找不到")
	case errors.As(err, &existErr):
		errRsp(c, errorType.CodeDataExist, "数据已经存在")
	case errors.As(err, &dbErr):
		errRsp(c, errorType.CodeDbError, "数据库错误")
	case errors.As(err, &logicError):
		errRsp(c, errorType.CodeLogicError, "逻辑错误")
	default:
		errRsp(c, errorType.CodeUnknown, "未知错误")
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
