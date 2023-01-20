package model

import "gorm.io/gorm"

type Entry struct {
	gorm.Model

	Title      string   `json:"title"`
	Amount     uint     `json:"amount"`
	TypeID     uint     `json:"typeID"`
	Type       Type     `json:"type"`
	CategoryID uint     `json:"categoryID"`
	Category   Category `json:"category"`
	UserID     uint     `json:"userID"`
}
