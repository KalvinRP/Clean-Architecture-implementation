package transactiondto

import models "dewetour/1models"

type TransactionResponse struct {
	ID       models.TransactionResponse `json:"user"`
	TotalPrc int                        `json:"totalprc"`
	Status   string                     `json:"status"`
	TripsID  int                        `json:"trips_id"`
	Trips    models.TripsResponse       `json:"trips" gorm:"foreignKey: TripsID"`
	UsersID  int                        `json:"users_id"`
	Users    models.UsersResponse       `json:"name" gorm:"foreignKey: UsersID"`
}
