package repositories

import (
	models "dewetour/1models"

	"gorm.io/gorm"
)

type TripsRepository interface {
	FindTrips() ([]models.Trips, error)
	GetTrips(ID int) (models.Trips, error)
	MakeTrips(trips models.Trips) (models.Trips, error)
	EditTrips(trips models.Trips, ID int) (models.Trips, error)
	DeleteTrips(trips models.Trips, ID int) (models.Trips, error)
}

func RepositoryTrips(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) MakeTrips(trips models.Trips) (models.Trips, error) {
	err := r.db.Create(&trips).Error

	return trips, err
}

func (r *repository) FindTrips() ([]models.Trips, error) {
	var trips []models.Trips
	err := r.db.Preload("User").Preload("Country").Find(&trips).Error

	return trips, err
}

func (r *repository) GetTrips(ID int) (models.Trips, error) {
	var trips models.Trips
	err := r.db.Preload("User").Preload("Country").First(&trips, ID).Error

	return trips, err
}

func (r *repository) EditTrips(trips models.Trips, ID int) (models.Trips, error) {
	var trip models.Trips
	err := r.db.Save(&trips).Error

	return trip, err
}

func (r *repository) DeleteTrips(trips models.Trips, ID int) (models.Trips, error) {
	var trip models.Trips
	err := r.db.Delete(&trip).Error

	return trip, err
}
