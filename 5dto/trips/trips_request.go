package tripsdto

type TripsRequest struct {
	Name           string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Desc           string `json:"desc" form:"desc" gorm:"type: text"`
	Price          string `json:"price" form:"price" gorm:"type: int"`
	Accomodation   string `json:"accomodation" form:"accomodation" gorm:"type: varchar(255)"`
	Transportation string `json:"transport" form:"transport" gorm:"type: varchar(255)"`
	Eat            string `json:"eat" form:"eat" gorm:"type: varchar(255)"`
	Duration       string `json:"duration" form:"duration" gorm:"type: varchar(255)"`
	DateTrip       string `json:"datetrip" form:"datetrip" gorm:"type: varchar(255)"`
	Qty            int    `json:"quota" form:"quota" gorm:"type: bit(64)"`
	Image          string `json:"image" form:"image" gorm:"type: varchar(255)"`
	UserID         int    `json:"user_id" gorm:"type: int"`
	CountryID      []int  `json:"country_id" form:"country_id" gorm:"type: int"`
}

type UpdateTripsRequest struct {
	Name           string `json:"name" form:"name"`
	Desc           string `json:"desc" form:"desc"`
	Price          string `json:"price" form:"price"`
	Accomodation   string `json:"accomodation" form:"accomodation"`
	Transportation string `json:"transport" form:"transport"`
	Eat            string `json:"eat" form:"eat"`
	Duration       string `json:"duration" form:"duration"`
	DateTrip       string `json:"datetrip" form:"datetrip"`
	Qty            int    `json:"quota" form:"quota"`
	Image          string `json:"image" form:"image"`
}
