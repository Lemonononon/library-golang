package service_book

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/define"
	"library/utils/response"
)

func CreateBook(c *gin.Context) response.Response {

	return response.Response{}
}

func GetBook(c *gin.Context, bookID int) response.Response {

	fmt.Printf("%d\n", bookID)
	return response.JSONStWithMsg(define.StOk, string(bookID))
}

func UpdateBook(c *gin.Context) response.Response {

	return response.Response{}
}

func DeleteBook(c *gin.Context) response.Response {

	return response.Response{}
}
