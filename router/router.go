package router

import (
	"github.com/gin-gonic/gin"
	"library/handler/handler_book"
	"library/handler/handler_record"
	"library/handler/ping"
	"library/mw"
)

func Register(r *gin.Engine) {
	r.Use(mw.CorsMiddleware) // CORS中间件  	必须在路由前配置

	api := r.Group("/api",
		mw.RecoverMiddleware,  // recover middleware
		mw.ResponseMiddleware, // response middleware
	)
	api.GET("/ping", ping.Pong) // ping
	book := api.Group("/book")
	{
		book.GET("/get_book", handler_book.GetBook)
		book.POST("/create_book", handler_book.CreateBook)
	}

	record := api.Group("/record")
	{
		record.POST("/borrow_book", handler_record.Borrow)
		record.POST("/return_book", handler_record.Return)
	}

}