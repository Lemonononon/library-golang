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
	c.Set(define.LibraryResponse, service_record.Borrow(c, card, bookID))
}

func Return(c *gin.Context) {
	cardID := c.Param("card_id")
	card := int(numberu.ToInt64(cardID))
	bookID := c.Param("book_id")
	c.Set(define.LibraryResponse, service_record.Return(c, card, bookID))
}

func QueryRecord(c *gin.Context) {
	cardID := c.Param("card_id")
	card := int(numberu.ToInt64(cardID))
	if card <= 0 {
		c.Set(define.LibraryResponse, define.StParamErr)
		return
	}

	c.Set(define.LibraryResponse, service_record.QueryRecord(c, card))
}
