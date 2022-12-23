package countrydto

type CountryRequest struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name"`
}

type UpdateCountryRequest struct {
	ID   int    `json:"id" form:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" form:"name"`
}
