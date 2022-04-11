package handler_auth

import (
	"github.com/gin-gonic/gin"
	"library/define"
	"library/model/model_auth"
	"library/service/service_auth"
)

func Login(c *gin.Context) {
	var req model_auth.LoginReq
	if err := c.ShouldBind(&req); err != nil {
		c.Set(define.LibraryResponse, define.StParamErr)
		return
	}
	c.Set(define.LibraryResponse, service_auth.Login(c, req))
}
