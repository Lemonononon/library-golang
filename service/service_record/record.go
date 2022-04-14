package service_record

import (
	"github.com/gin-gonic/gin"
	"library/db"
	"library/define"
	"library/model/model_record"
	"library/utils/response"
	"library/utils/timeu"
	"time"
)

func Borrow(c *gin.Context, cardID int, bookID string) response.Response {
	record := model_record.Record{
		CardID:     cardID,
		BookID:     bookID,
		BorrowDate: timeu.FormatDate(time.Now()),
	}
	if err := db.MySQL.Create(&record).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}

	res := model_record.BorrowResp{
		Done: "Success!",
	}
	return response.JSONData(res)
}

func Return(c *gin.Context, cardID int, bookID string) response.Response {
	var record model_record.Record
	if err := db.MySQL.Where("card_id = ? AND book_id = ?", cardID, bookID).First(&record).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}

	if err := db.MySQL.Delete(&record).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}
	return response.JSONData("Succeed!")
}

func QueryRecord(c *gin.Context, cardID int) response.Response {
	var records model_record.RecordQueryResp

	if err := db.MySQL.Where("card_id = ?", cardID).Find(&records.Records).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}

	return response.JSONData(records)
}
