package service_book

import (
	"github.com/gin-gonic/gin"
	"library/db"
	"library/define"
	"library/model/model_book"
	"library/utils/response"
)

func QueryBooks(c *gin.Context) response.Response {
	return response.Response{}
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

func DeleteBook(c *gin.Context) response.Response {

	return response.Response{}
}
