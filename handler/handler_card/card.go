package handler_card

import (
	"github.com/gin-gonic/gin"
	"library/define"
	"library/model/model_card"
	"library/service/service_card"
	"library/utils/numberu"
)

func AddCard(c *gin.Context) {
	var req model_card.AddCardReq
	if err := c.ShouldBind(&req); err != nil {
		c.Set(define.LibraryResponse, define.StParamErr)
		return
	}
	c.Set(define.LibraryResponse, service_card.AddCard(c, req))
}

func DeleteCard(c *gin.Context) {
	tmpID := c.Param("card_id")
	cardID := int(numberu.ToInt64(tmpID))
	if cardID <= 0 {
		c.Set(define.LibraryResponse, define.StNoCard)
		return
	}
	c.Set(define.LibraryResponse, service_card.DeleteCard(c, cardID))
}
