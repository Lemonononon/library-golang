package service_record

import (
	"github.com/gin-gonic/gin"
	"library/db"
	"library/define"
	"library/model/model_book"
	"library/model/model_record"
	"library/utils/response"
	"library/utils/timeu"
	"time"
)

func Borrow(c *gin.Context, cardID int, bookID string) response.Response {
	var book model_book.Book
	if err := db.MySQL.Where("book_id = ?", bookID).First(&book).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}
	if book.Stock <= 0 {
		return response.JSONSt(define.StNoEnoughBook)
	}
	record := model_record.Record{
		Status:     0,
		CardID:     cardID,
		BookID:     bookID,
		BorrowDate: timeu.FormatDate(time.Now()),
	}
	if err := db.MySQL.Create(&record).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}
	book.Stock--
	if err := db.MySQL.Save(&book).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}
	res := model_record.BorrowResp{
		Done: "Success!",
	}
	return response.JSONData(res)
}

func Return(c *gin.Context, cardID int, bookID string) response.Response {
	var record model_record.Record
	if err := db.MySQL.Where("card_id = ? AND book_id = ? AND status = ?", cardID, bookID, 0).Last(&record).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}

	record.Status = 1 // 已归还
	record.ReturnDate = timeu.FormatDate(time.Now())
	if err := db.MySQL.Save(&record).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}
	//if err := db.MySQL.Delete(&record).Error; err != nil {
	//	return response.JSONSt(define.StDBErr)
	//}
	var book model_book.Book
	if err := db.MySQL.Where("book_id = ?", bookID).First(&book).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}
	book.Stock++
	if err := db.MySQL.Save(&book).Error; err != nil {
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
