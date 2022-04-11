package handler_auth

import (
	"github.com/gin-gonic/gin"
	"library/define"
)

func Login(c *gin.Context) {

	c.Set(define.LibraryResponse, service)
}