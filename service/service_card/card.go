package service_card

import (
	"github.com/gin-gonic/gin"
	"library/db"
	"library/define"
	"library/model/model_card"
	"library/utils/response"
)

func AddCard(c *gin.Context, req model_card.AddCardReq) response.Response {

	card := model_card.Card{
		Name:       req.Name,
		Department: req.Department,
		Identity:   req.Identity,
	}
	if err := db.MySQL.Create(&card).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}
	res := model_card.AddCardResp{
		CardID: card.CardID,
	}
	return response.JSONData(res)
}

func DeleteCard(c *gin.Context, cardID int) response.Response {
	var card model_card.Card
	if err := db.MySQL.Delete(&card, "card_id = ?", cardID).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}
	return response.JSONData(model_card.DeleteCardReq{
		Done: "Succeed!",
	})
}
