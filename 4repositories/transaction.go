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
	err := r.db.Preload("Trips").Preload("Users").Find(&transaction).Error

	return transaction, err
}

func (r *repository) GetTrans(ID int) (models.Transaction, error) {
	var trans models.Transaction
	err := r.db.Preload("Trips").Preload("Users").First(&trans, ID).Error

	return trans, err
}

func (r *repository) MakeTrans(trans models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&trans).Error

	return trans, err
}

func (r *repository) EditTrans(trans models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Model(&trans).Updates(trans).Error

	return trans, err
}

func (r *repository) DeleteTrans(trans models.Transaction, ID int) (models.Transaction, error) {
	err := r.db.Delete(&trans).Error

	return trans, err
}
