package models

import "time"

type Transaction struct {
	ID        int           `json:"id" gorm:"primary_key:auto_increment"`
	TripsID   int           `json:"trips_id"`
	Trips     TripsResponse `json:"trips" gorm:"foreignKey:TripsID"`
	UsersID   int           `json:"users_id"`
	Users     UsersResponse `json:"name" gorm:"foreignKey: UsersID"`
	UserID    int           `json:"user_id" gorm:"type: int"`
	TotalPrc  int           `json:"totalprc"`
	Status    string        `json:"status" gorm:"type: varchar(20)"`
	CreatedAt time.Time     `json:"-"`
	UpdatedAt time.Time     `json:"-"`
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
