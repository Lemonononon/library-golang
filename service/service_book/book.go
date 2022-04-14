package service_book

import (
	"github.com/gin-gonic/gin"
	"library/db"
	"library/define"
	"library/model/model_book"
	"library/utils/response"
)

func QueryBooks(c *gin.Context, req model_book.QueryBookReq) response.Response {
	tx := db.MySQL.Begin()
	lowerYear, upperYear := req.LowerYear, req.UpperYear
	if lowerYear != 0 && upperYear != 0 {
		tx = tx.Where("year >= ? and year <= ?", lowerYear, upperYear)
	}
	lowerPrice, upperPrice := req.LowerPrice, req.UpperPrice
	if lowerPrice != 0 || upperPrice != 0 {
		tx = tx.Where("price >= ?", lowerPrice).Where("price <= ?", upperPrice)
	}

	book := &model_book.Book{
		BookName: req.BookName,
		Category: req.Category,
		Press:    req.Press,
		Author:   req.Author,
	}
	var books []model_book.Book
	//var books []model_book.Book
	//if req.LowerPrice != {}
	if err := tx.Debug().Order(req.Order).Where(&book).Find(&books).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}
	return response.JSONData(&model_book.QueryBookResp{Books: books})
}

func AddBook(c *gin.Context, req model_book.AddBookReq) response.Response {
	book := model_book.Book{
		BookID:   req.BookID,
		Category: req.Category,
		BookName: req.BookName,
		Press:    req.Press,
		Year:     req.Year,
		Author:   req.Author,
		Price:    req.Price,
		Total:    req.Number,
		Stock:    req.Number,
	}
	if err := db.MySQL.Create(&book).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}

	res := model_book.AddBookResp{
		BookID: book.BookID,
	}
	return response.JSONData(res)
}

func AddBooks(c *gin.Context, req model_book.AddBooksReq) response.Response {
	var books []model_book.Book

	for _, v := range req.Books {
		books = append(books, model_book.Book{
			BookID:   v.BookID,
			Category: v.Category,
			BookName: v.BookName,
			Press:    v.Press,
			Year:     v.Year,
			Author:   v.Author,
			Price:    v.Price,
			Total:    v.Number,
			Stock:    v.Number,
		})
	}

	if err := db.MySQL.Create(&books).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}
	res := model_book.AddBooksResp{
		Done: "All done!",
	}
	return response.JSONData(res)
}

func DeleteBook(c *gin.Context, bookID string) response.Response {
	var book model_book.Book
	if err := db.MySQL.Where("book_id=?", bookID).First(&book).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}
	if err := db.MySQL.Delete(&book).Error; err != nil {
		return response.JSONSt(define.StDBErr)
	}
	res := model_book.DeleteBookResp{
		Done: "Delete Succeed!",
	}
	return response.JSONData(res)
}
