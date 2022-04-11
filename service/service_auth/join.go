package service_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/auth"
	"library/db"
	"library/define"
	"library/model/model_admin"
	"library/model/model_auth"
	"library/service/service_admin"
	"library/utils/crypto"
	"library/utils/response"
)

func Join(c *gin.Context, req model_auth.JoinReq) response.Response {

	admin := &model_admin.Admin{

		Password:  crypto.Password2Secret(req.Password), // 密码加密
		AdminName: req.AdminName,
		Phone:     req.Phone,
	}

	//进行注册前的check
	if !service_admin.ValidAdmin(admin) {
		return response.JSONSt(define.StParamErr)
	}
	// 操作数据库
	result := db.MySQL.Create(admin)

	if result.Error != nil {
		return response.JSONSt(define.StDuplicateJoinErr)
	}
	//生成token
	token, err := auth.GenToken(admin.AdminID, auth.RoleTypeUser)
	if err != nil {
		fmt.Println(err)
		return response.JSONSt(define.StServerErr)
	}

	res := model_auth.JoinResp{
		Token: token,
	}
	return response.JSONData(res)
}
