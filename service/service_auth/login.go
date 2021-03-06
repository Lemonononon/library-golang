package service_auth

import (
	"github.com/gin-gonic/gin"
	"library/auth"
	"library/define"
	"library/model/model_auth"
	"library/service/service_admin"
	"library/utils/crypto"
	"library/utils/response"
)

func Login(c *gin.Context, req model_auth.LoginReq) response.Response {

	//adminID := int(numberu.ToInt64(req.AdminID))
	admin := service_admin.QueryAdminByPhone(req.Phone)
	if admin.AdminID < 0 {
		return response.JSONSt(define.StNoUser)
	}
	if admin.Phone == req.Phone && admin.Password == crypto.Password2Secret(req.Password) {
		token, err := auth.GenToken(admin.AdminID, admin.Role)
		if err != nil {
			return response.JSONSt(define.StServerErr)
		}
		res := model_auth.LoginResp{
			Token: token,
		}
		return response.JSONData(res)
	}

	return response.JSONSt(define.StLoginErr)
}
