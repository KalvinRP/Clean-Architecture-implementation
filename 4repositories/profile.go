package repositories

import (
	models "dewetour/1models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	FindProfile() ([]models.Profile, error)
	GetProfile(ID int) (models.Profile, error)
	MakeProfile(profile models.Profile) (models.Profile, error)
	EditProfile(profile models.Profile, ID int) (models.Profile, error)
	DeleteProfile(profile models.Profile, ID int) (models.Profile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProfile() ([]models.Profile, error) {
	var profile []models.Profile
	err := r.db.Preload("Profile").Preload("Trips").Find(&profile).Error

	return profile, err
}

func (r *repository) GetProfile(ID int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Preload("User").First(&profile, ID).Error

	return profile, err
}

func (r *repository) MakeProfile(user models.Profile) (models.Profile, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) EditProfile(user models.Profile, ID int) (models.Profile, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteProfile(user models.Profile, ID int) (models.Profile, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
