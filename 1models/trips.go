package models

import "time"

type Trips struct {
	ID             int                  `json:"id" gorm:"primary_key:auto_increment"`
	Name           string               `json:"name" form:"name" gorm:"type: varchar(255)"`
	Desc           string               `json:"desc" form:"desc" gorm:"type: text"`
	Price          string               `json:"price" form:"price" gorm:"type: int"`
	Accomodation   string               `json:"accomodation" form:"accomodation" gorm:"type: varchar(255)"`
	Transportation string               `json:"transport" form:"transport" gorm:"type: varchar(255)"`
	Eat            string               `json:"eat" form:"eat" gorm:"type: varchar(255)"`
	Duration       string               `json:"duration" form:"duration" gorm:"type: varchar(255)"`
	DateTrip       string               `json:"datetrip" form:"datetrip" gorm:"type: varchar(255)"`
	Qty            int                  `json:"quota" form:"quota" gorm:"type: bit(64)"`
	Image          string               `json:"image" form:"image" gorm:"type: varchar(255)"`
	User           UsersProfileResponse `json:"user"`
	UserID         int                  `json:"user_id" form:"user_id"`
	Country        []Country            `json:"country" gorm:"many2many:trip_countries"`
	CountryID      []int                `json:"country_id" form:"country_id" gorm:"-"`
	CreatedAt      time.Time            `json:"-"`
	UpdatedAt      time.Time            `json:"-"`
}

type TripsResponse struct {
	ID             int                  `json:"id"`
	Name           string               `json:"name"`
	Desc           string               `json:"desc"`
	Price          string               `json:"price"`
	Accomodation   string               `json:"accomodation"`
	Transportation string               `json:"transport"`
	Eat            string               `json:"eat"`
	Duration       string               `json:"duration"`
	DateTrip       string               `json:"datetrip"`
	Qty            int                  `json:"quota"`
	Image          string               `json:"image"`
	User           UsersProfileResponse `json:"user"`
	UserID         int                  `json:"user_id"`
	Country        []Country            `json:"country" gorm:"many2many:trip_countries"`
	CountryID      []int                `json:"country_id" form:"country_id" gorm:"-"`
}

func (TripsResponse) TableName() string {
	return "trips"
}

type TripsUserResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Desc           string `json:"desc"`
	Price          string `json:"price"`
	Accomodation   string `json:"accomodation"`
	Transportation string `json:"transport"`
	Eat            string `json:"eat"`
	Duration       string `json:"duration"`
	DateTrip       string `json:"datetrip"`
	Qty            int    `json:"quota"`
	Image          string `json:"image"`
	UserID         int    `json:"-"`
}

func (TripsUserResponse) TableName() string {
	return "trips"
}
