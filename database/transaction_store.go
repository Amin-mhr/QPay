package database

import (
	"qpay/models"

	"gorm.io/gorm"
)

func PostTransaction(transaction models.Transaction, db *gorm.DB) error {
	err := db.Create(&transaction)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
