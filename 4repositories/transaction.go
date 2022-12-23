package repositories

import (
	models "dewetour/1models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTrans() ([]models.Transaction, error)
	GetTrans(ID int) (models.Transaction, error)
	MakeTrans(transaction models.Transaction) (models.Transaction, error)
	EditTrans(transaction models.Transaction, ID int) (models.Transaction, error)
	DeleteTrans(transaction models.Transaction, ID int) (models.Transaction, error)
}

func RepositoryTrans(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTrans() ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("Profile").Preload("Trips").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTrans(ID int) (models.Transaction, error) {
	var user models.Transaction
	err := r.db.Preload("Profile").Preload("Trips").First(&user, ID).Error

	return user, err
}

func (r *repository) MakeTrans(user models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) EditTrans(user models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteTrans(user models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
