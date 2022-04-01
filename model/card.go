package model

type Card struct {
	CardID     int    `gorm:"column:card_id;primaryKey"`
	Name       string `gorm:"column:name;size:20"` //TODO 外键
	Department string `gorm:"column:department;size:40"`
	Category   string `gorm:"column:category;size:20"`
}
