package service_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/auth"
	"library/define"
	"library/model/model_admin"
	"library/model/model_auth"
	"library/utils/response"
)

func Join(c *gin.Context, req model_auth.JoinReq) response.Response {

	user := &model_admin.Admin{

		Password:  crypto.Password2Secret(req.Password),
		AdminName: req.AdminName,
		Phone:     req.Phone,
	}

	//进行用户注册前的check
	if !service_user.ValidUser(user) {
		return response.JSONSt(define.StParamErr)
	}

	//加密密码
	user.Password = crypto.Password2Secret(user.Password)

	// 操作数据库
	result := caller.MySqlDB.Create(user)

	if result.Error != nil {
		return response.JSONSt(define.StDuplicateJoinErr)
	}
	//生成token
	token, err := auth.GenToken(user.ID, auth.RoleTypeUser)
	if err != nil {
		fmt.Println(err)
		return response.JSONSt(define.StServerErr)
	}

	res := model_auth.JoinResp{
		Token: token,
	}
	return response.JSONData(res)
}
