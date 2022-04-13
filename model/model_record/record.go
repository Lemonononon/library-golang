package model_record

type Record struct {
	ID         int    `gorm:"column:id""`
	BookID     int    `gorm:"column:book_id"`
	CardID     int    `gorm:"column:card_id"`
	BorrowDate string `gorm:"column:borrow_date"`
	ReturnDate string `gorm:"column:return_date"`
}

type RecordQueryResp struct {
	Records []Record `json:"records"`
}

type BorrowReq struct {
	BookID     int    `json:"book_id" gorm:"column:book_id"`
	CardID     int    `json:"card_id" gorm:"column:card_id"`
	BorrowDate string `json:"borrow_date" gorm:"column:borrow_date"`
}
type BorrowResp struct {
	Done string `json:"done"`
}
