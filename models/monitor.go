package models

import (
	"time"
)

type Monitor struct {
	ID              uint `gorm:"primaryKey"`
	Url             string
	Title           string
	HashValuesTitle string    `gorm:"column:hash_values_title"`
	HashValuesBody  string    `gorm:"column:hash_values_body"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}

type MonitorView struct {
	Time        string
	TitleChange float64
	BodyChange  float64
}
