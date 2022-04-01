package main

import (
	"github.com/gin-gonic/gin"
	"library/handler/ping"
	"library/mw"
)

func register(r *gin.Engine) {
	r.Use(mw.CorsMiddleware) // CORS中间件  	必须在路由前配置

	api := r.Group("/api",
		mw.RecoverMiddleware,  // recover middleware
		mw.ResponseMiddleware, // response middleware
	)
	api.GET("/ping", ping.Pong) // ping
}
