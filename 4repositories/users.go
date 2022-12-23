package repositories

import (
	models "dewetour/1models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAcc() ([]models.User, error)
	GetAcc(ID int) (models.User, error)
	MakeAcc(user models.User) (models.User, error)
	EditAcc(user models.User, ID int) (models.User, error)
	DeleteAcc(user models.User, ID int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAcc() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Profile").Preload("Trips").Find(&users).Error

	return users, err
}

func (r *repository) GetAcc(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Profile").Preload("Trips").First(&user, ID).Error

	return user, err
}

func (r *repository) MakeAcc(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) EditAcc(user models.User, ID int) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteAcc(user models.User, ID int) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
