package mw

import (
	"github.com/gin-gonic/gin"
	"library/define"
	"library/utils/response"
	"log"
)

func ResponseMiddleware(c *gin.Context) {
	c.Next()
	value, exists := c.Get(define.LibraryResponse)
	if !exists {
		log.Println("[ResponseMiddleware] response not set!")
		return
	}
	resp, ok := value.(response.Response)
	if !ok {
		log.Println("[ResponseMiddleware] response type invalid!")
		return
	}
	resp.Write(c)
}
