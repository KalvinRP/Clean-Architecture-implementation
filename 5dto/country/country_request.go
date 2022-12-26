package countrydto

type CountryRequest struct {
	ID     int    `json:"id" gorm:"primary_key:auto_increment"`
	Name   string `json:"name"`
	UserID int    `json:"user_id" gorm:"type: int"`
}
