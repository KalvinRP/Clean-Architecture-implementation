package authdto

type AuthRequest struct {
	Name     string `json:"name" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Address  string `json:"address" form:"address" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
}
