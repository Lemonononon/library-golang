package model_card

import "library/define"

type Card struct {
	CardID     int                 `gorm:"column:card_id;primaryKey"`
	Name       string              `gorm:"column:name;size:20"` //TODO 外键
	Department string              `gorm:"column:department;size:40"`
	Identity   define.CardIdentity `gorm:"column:identity"` //0-teacher 1-student
}

type AddCardReq struct {
	Name       string              `json:"name"`
	Department string              `json:"department"`
	Identity   define.CardIdentity `json:"identity"`
}

type AddCardResp struct {
	CardID int `json:"card_id"`
}
