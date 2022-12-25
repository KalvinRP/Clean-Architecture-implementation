package models

import "time"

type Transaction struct {
	ID      int           `json:"id" gorm:"primary_key:auto_increment"`
	TripsID int           `json:"trips_id"`
	Trips   TripsResponse `json:"accomodation" gorm:"foreignKey: TripsID"`
	UsersID int           `json:"users_id"`
	Users   UsersResponse `json:"name" gorm:"foreignKey: UsersID"`
	// Name           Trips         `json:"trips"`
	// Desc           Trips         `json:"desc"`
	// Price          Trips         `json:"price"`
	// Eat            Trips         `json:"eat"`
	// Quota          Trips         `json:"quota"`
	// Image          Trips         `json:"image"`
	// Country        Trips         `json:"country"`
	TotalPrc  int       `json:"totalprc"`
	Status    string    `json:"status" gorm:"type: varchar(20)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type TransactionResponse struct {
	ID       int           `json:"id"`
	TotalPrc string        `json:"totalprc"`
	Status   string        `json:"status"`
	TripsID  int           `json:"trips_id"`
	Trips    Trips         `json:"accomodation" gorm:"foreignKey: TripsID"`
	UsersID  int           `json:"users_id"`
	Users    UsersResponse `json:"name" gorm:"foreignKey: UsersID"`
}

func (TransactionResponse) TableName() string {
	return "profiles"
}
