package model_record

type Record struct {
	BookID     int    `gorm:""` //TODO: 外键
	CardID     int    `gorm:""`
	BorrowDate string `gorm:""`
	ReturnDate string `gorm:""`
}
