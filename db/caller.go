package db

import "gorm.io/gorm"

var MySQL *gorm.DB

func InitCaller() {
	initMySQL()
}
