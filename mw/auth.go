package mw

import (
	"github.com/gin-gonic/gin"
	"library/auth"
	"library/define"
	"library/utils/response"
	"strings"
	"time"
)

// AuthenticationMiddleware 身份验证
func AuthenticationMiddleware(c *gin.Context) {
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.Set(define.LibraryResponse, response.JSONSt(define.StNoLoginUser))
		c.Abort()
		return
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.Set(define.LibraryResponse, response.JSONSt(define.StParamErr))
		c.Abort()
		return
	}
	// parts[1]是获取到的tokenString
	mc, err := auth.ParseToken(parts[1])
	if err != nil {
		c.Set(define.LibraryResponse, response.JSONSt(define.StParamErr))
		c.Abort()
		return
	} else if mc.ExpiresAt <= time.Now().Unix() {
		c.Set(define.LibraryResponse, response.JSONSt(define.StTokenExpired))
		c.Abort()
		return
	}
	// 将当前请求的信息保存到请求的上下文c上
	c.Set(define.LibraryUserID, mc.ID)
	c.Set(define.LibraryUserRole, mc.Role)
	c.Next()
}

// AuthorizationMiddleware 授权
func AuthorizationMiddleware(roles []auth.RoleType) func(c *gin.Context) {
	return func(c *gin.Context) {
		userRole, ok := c.Value(define.LibraryUserRole).(auth.RoleType)
		if !ok {
			c.Set(define.LibraryResponse, response.JSONSt(define.StPermissionErr))
			c.Abort()
			return
		}
		if !isContain(roles, userRole) {
			c.Set(define.LibraryResponse, response.JSONSt(define.StPermissionErr))
			c.Abort()
			return
		}
		c.Next()
	}

}

func isContain(items []auth.RoleType, item auth.RoleType) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
