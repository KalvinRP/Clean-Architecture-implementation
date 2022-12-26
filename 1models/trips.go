package models

import "time"

type Trips struct {
	ID             int       `json:"id" gorm:"primary_key:auto_increment"`
	Name           string    `json:"name" form:"name" gorm:"type: varchar(255)"`
	Desc           string    `json:"desc" form:"desc" gorm:"type: text"`
	Price          int       `json:"price" form:"price" gorm:"type: int"`
	Accomodation   string    `json:"accomodation" form:"accomodation" gorm:"type: varchar(255)"`
	Transportation string    `json:"transport" form:"transport" gorm:"type: varchar(255)"`
	Eat            string    `json:"eat" form:"eat" gorm:"type: varchar(255)"`
	Duration       string    `json:"duration" form:"duration" gorm:"type: varchar(255)"`
	DateTrip       string    `json:"datetrip" form:"datetrip" gorm:"type: varchar(255)"`
	Quota          int       `json:"quota" form:"quota" gorm:"type: int"`
	Image          string    `json:"image" form:"image" gorm:"type: varchar(255)"`
	Country        Country   `json:"country"`
	CountryID      int       `json:"country_id" form:"country_id"`
	UserID         int       `json:"user_id" form:"user_id"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}

type TripsResponse struct {
	ID             int    `json:"id"`
	Accomodation   string `json:"accomodation"`
	Transportation string `json:"transport"`
	Duration       string `json:"duration"`
	DateTrip       string `json:"datetrip"`
}

func (TripsResponse) TableName() string {
	return "trips"
}
