package model_admin

import "library/auth"

type Admin struct {
	AdminID   int           `gorm:"column:administrator_id;primaryKey"`
	Password  string        `gorm:"column:password"`
	AdminName string        `gorm:"column:admin_name"`
	Role      auth.RoleType `gorm:"column:role"`
	Phone     string        `gorm:"column:phone;uniqueIndex;size:15"`
}
