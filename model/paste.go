package model

import (
	"time"
)

type PasteBase struct {
	Title string `json:"title"`
	Content string `json:"content"`
	ContentType string `json:"content-type"`
}

type PasteModel struct {
	ID string `gorm:"primary_key"`
	PasteBase
	Timestamp time.Time
}

func (PasteModel) TableName() string {
	return "pastes"
}
