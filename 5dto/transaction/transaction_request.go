package transactiondto

type TransactionRequest struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	TripsID  int    `json:"trips_id"`
	UsersID  int    `json:"Users_id"`
	TotalPrc int    `json:"totalprc"`
	Status   string `json:"status" gorm:"type: varchar(20)"`
}

type UpdateTransactionRequest struct {
	ID       int    `json:"id" form:"id"`
	TripsID  int    `json:"tripsid" form:"tripsid"`
	UsersID  int    `json:"Usersid" form:"Usersid"`
	TotalPrc int    `json:"totalprc" form:"totalprc"`
	Status   string `json:"status"`
}
