package router

import (
	"github.com/gin-gonic/gin"
	"library/handler/handler_auth"
	"library/handler/handler_book"
	"library/handler/handler_card"
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

	api.POST("/login", handler_auth.Login) // 登录
	api.POST("/join", handler_auth.Join)   // 注册
	//api.Use(mw.AuthenticationMiddleware)   // 开启身份验证

	book := api.Group("/book")
	{
		book.POST("/add", handler_book.AddBook)      // 添加一本书
		book.POST("/adds", handler_book.AddBooks)    // 批量添加书籍
		book.POST("/query", handler_book.QueryBooks) // 查询书籍
	}
	card := api.Group("/card")
	{
		card.POST("/add", handler_card.AddCard)
		card.DELETE("/delete/:card_id", handler_card.DeleteCard)
	}
	record := api.Group("/record")
	{
		record.GET("/query/:card_id", handler_record.QueryRecord)
		record.GET("/borrow/:card_id/:book_id", handler_record.Borrow)
		record.GET("/return", handler_record.Return)
	}
}
