package model_admin

type Admin struct {
	AdminID   int    `gorm:"column:administrator_id;primaryKey"`
	Password  string `gorm:"column:password"`
	AdminName string `gorm:"column:admin_name"`
	Phone     string `gorm:"column:phone;uniqueIndex;size:15"`
}
