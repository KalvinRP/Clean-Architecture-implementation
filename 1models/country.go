package models

import "time"

type Country struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name"`
	UserID    int       `json:"user_id" gorm:"type: int"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type CountryResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserID int    `json:"user_id" gorm:"type: int"`
}

func (CountryResponse) TableName() string {
	return "countries"
}
