package authdto

type LoginResponse struct {
	Name  string `gorm:"type: varchar(255)" json:"name"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}
