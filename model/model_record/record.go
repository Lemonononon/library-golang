package model_record

type Record struct {
	ID         int    `gorm:"column:id"`
	BookID     string `gorm:"column:book_id"`
	CardID     int    `gorm:"column:card_id"`
	Status     int    `gorm:"column:status"` //0:未归还 1:已归还
	BorrowDate string `gorm:"column:borrow_date"`
	ReturnDate string `gorm:"column:return_date"`
}

type RecordQueryResp struct {
	Records []Record `json:"records"`
}

type BorrowReq struct {
	BookID int `json:"book_id" gorm:"column:book_id"`
	CardID int `json:"card_id" gorm:"column:card_id"`
}
type BorrowResp struct {
	Done string `json:"done"`
}
