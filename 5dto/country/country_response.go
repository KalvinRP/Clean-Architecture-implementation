package countrydto

import models "dewetour/1models"

type CountryResponse struct {
	ID   int                    `json:"id" gorm:"primary_key:auto_increment"`
	Name models.CountryResponse `json:"user"`
}
