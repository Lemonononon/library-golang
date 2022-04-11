package handler_auth

import (
	"github.com/gin-gonic/gin"
	"library/define"
	"library/model/model_auth"
	"library/service/service_auth"
)

func Join(c *gin.Context) {
	var req model_auth.JoinReq
	//tmpAdminID, _ := c.Get("admin_id")
	//AdminID, _ := tmpAdminID.(int)
	//if AdminID <= 0 {
	//	c.Set(define.LibraryResponse, define.StNoUser)
	//}
	//
	if err := c.ShouldBind(&req); err != nil {
		c.Set(define.LibraryResponse, define.StParamErr)
		return
	}
	c.Set(define.LibraryResponse, service_auth.Join(c, req))
}
