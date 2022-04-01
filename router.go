package main

import (
	"github.com/gin-gonic/gin"
	"library/mw"
)

func register(r *gin.Engine) {
	r.Use(mw.CorsMiddleware) // CORS中间件  	必须在路由前配置
}
