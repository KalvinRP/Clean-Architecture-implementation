package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"password" gorm:"type: varchar(255)"`
	Phone     string    `json:"phone" gorm:"type: varchar(20)"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	Gender    string    `json:"gender" gorm:"type: varchar(20)"`
	Role      string    `gorm:"type:varchar(5)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (UsersResponse) TableName() string {
	return "users"
}
