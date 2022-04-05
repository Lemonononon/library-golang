package handler_book

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/define"
	"library/service/service_book"
	"library/utils/numberu"
)

func CreateBook(c *gin.Context) {

	c.Set(define.LibraryResponse, service_book.CreateBook(c))
}
func GetBook(c *gin.Context) {
	bookID := c.Query("book_id")
	//bookID := c.Param("book_id")
	book := int(numberu.ToInt64(bookID))
	fmt.Printf("%s\n", bookID)
	c.Set(define.LibraryResponse, service_book.GetBook(c, book))
}
