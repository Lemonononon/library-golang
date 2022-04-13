package ping

import (
	"github.com/gin-gonic/gin"
	"library/define"
	"library/utils/response"
)

func Pong(c *gin.Context) {
	c.Set(define.LibraryResponse, response.JSONData("pong!"))
}
