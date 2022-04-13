package service_record

import (
	"github.com/gin-gonic/gin"
	"library/db"
	"library/define"
	"library/model/model_record"
	"library/utils/response"
)

func Borrow(c *gin.Context, cardID int, bookID string) response.Response {
	var record model_record.Record

	if err := db.MySQL.Create(&record).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}

	res := model_record.BorrowResp{
		Done: "Success!",
	}
	return response.JSONData(res)
}

func Return(c *gin.Context, cardID int, bookID string) response.Response {

	return response.Response{}
}

func QueryRecord(c *gin.Context, cardID int) response.Response {
	var records model_record.RecordQueryResp

	if err := db.MySQL.Where("card_id = ?", cardID).Find(&records).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}

	return response.JSONData(records)
}
