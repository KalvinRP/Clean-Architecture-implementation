package transactiondto

import models "dewetour/1models"

type TransactionResponse struct {
	ID       models.TransactionResponse `json:"user"`
	BookerID int                        `json:"-"`
	Trips    models.TripsResponse       `json:"trips"`
	TotalPrc int                        `json:"totalprc"`
	Status   string                     `json:"status"`
}
