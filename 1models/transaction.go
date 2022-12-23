package models

import "time"

type Transaction struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	TripsID   int                  `json:"trips_id"`
	Trips     TripsResponse        `json:"trips"`
	BookerID  int                  `json:"booker_id"`
	Booker    UsersProfileResponse `json:"booker"`
	TotalPrc  int                  `json:"totalprc"`
	Status    string               `json:"status" gorm:"type: varchar(20)"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

type TransactionResponse struct {
	ID       int           `json:"id"`
	BookerID int           `json:"-"`
	Trips    TripsResponse `json:"trips"`
	TotalPrc string        `json:"totalprc"`
	Status   string        `json:"status"`
}

func (TransactionResponse) TableName() string {
	return "profiles"
}
