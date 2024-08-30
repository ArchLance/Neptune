package controller

import (
	"github.com/gin-gonic/gin"
	"neptune/utils/rsp"
	jwt "neptune/utils/token"
)

func Login(c *gin.Context) {
	//Login := forms.Login{}
	//if err := c.ShouldBind(&Login); err != nil {
	//	log.Errorf("Login ShouldBind Error: %v", err)
	//	// 统一处理异常
	//	util.HandleValidatorError(c, err)
	//	return
	//}
	//user, ok, _ := dao.FindUserInfo(Login.Telephone, Login.Password)
	//if !ok {
	//	response.Err(c, http.StatusUnauthorized, 401, "用户未注册或密码错误", nil)
	//	return
	//}
	userID := 10000
	userRole := "admin"
	token := jwt.GenerateToken(c, userID, userRole)
	rsp.SuccessRsp(c, token)
}
