package model

import "time"

type ShortURL struct {
	ID         string `gorm:"primary_key"`
	ShortUrl   string
	LangUrl    string
	Duration   string
	Token      string
	CreateTime time.Time
	UpdateTime time.Time
}

func (s ShortURL) TableName() string {
	return "short_url"
}
