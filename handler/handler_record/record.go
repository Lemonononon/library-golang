package handler_record

import (
	"github.com/gin-gonic/gin"
	"library/define"
	"library/service/service_record"
	"library/utils/numberu"
)

func Borrow(c *gin.Context) {
	cardID := c.Param("card_id")
	card := int(numberu.ToInt64(cardID))
	bookID := c.Param("book_id")
	book := int(numberu.ToInt64(bookID))
	c.Set(define.LibraryResponse, service_record.BorrowBook(c, card, book))
}

func Return(c *gin.Context) {
	cardID := c.Param("card_id")
	card := int(numberu.ToInt64(cardID))
	bookID := c.Param("book_id")
	book := int(numberu.ToInt64(bookID))
	c.Set(define.LibraryResponse, service_record.ReturnBook(c, card, book))
}
