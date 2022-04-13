package handler_book

import (
	"github.com/gin-gonic/gin"
	"library/define"
	"library/service/service_book"
)

//func CreateBook(c *gin.Context) {
//
//	c.Set(define.LibraryResponse, service_book.CreateBook(c))
//}
//func GetAllBook(c *gin.Context) {
//	bookID := c.Query("book_id")
//	//bookID := c.Param("book_id")
//	book := int(numberu.ToInt64(bookID))
//	c.Set(define.LibraryResponse, service_book.GetBook(c, book))
//}

func QueryBooks(c *gin.Context) {
	c.Set(define.LibraryResponse, service_book.QueryBooks(c))
}

func AddBook(c *gin.Context) {

}
func AddBooks(c *gin.Context) {

}
