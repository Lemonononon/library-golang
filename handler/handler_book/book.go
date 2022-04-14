package handler_book

import (
	"github.com/gin-gonic/gin"
	"library/define"
	"library/model/model_book"
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
	var req model_book.QueryBookReq
	if err := c.ShouldBind(&req); err != nil {
		c.Set(define.LibraryResponse, define.StParamErr)
		return
	}
	c.Set(define.LibraryResponse, service_book.QueryBooks(c, req))
}

func AddBook(c *gin.Context) {
	var req model_book.AddBookReq
	if err := c.ShouldBind(&req); err != nil {
		c.Set(define.LibraryResponse, define.StParamErr)
		return
	}

	c.Set(define.LibraryResponse, service_book.AddBook(c, req))
}

//AddBooks book_no_1, Computer Science, Computer Architecture, xxx, 2004, xxx, 90.00, 2
func AddBooks(c *gin.Context) {
	var req model_book.AddBooksReq
	if err := c.ShouldBind(&req); err != nil {
		c.Set(define.LibraryResponse, define.StParamErr)
		return
	}
	c.Set(define.LibraryResponse, service_book.AddBooks(c, req))
}

func DeleteBook(c *gin.Context) {
	bookID := c.Param("book_id")
	if len(bookID) == 0 {
		c.Set(define.LibraryResponse, define.StParamErr)
		return
	}
	c.Set(define.LibraryResponse, service_book.DeleteBook(c, bookID))
}
