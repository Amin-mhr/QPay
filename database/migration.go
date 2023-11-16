package database

import (
	"gorm.io/gorm"
	models "signup/models"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&models.Admin{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.Pricing{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.Account{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.Gateway{})
	if err != nil {
		return
	}

	err = db.AutoMigrate(&models.Transaction{})
	if err != nil {
		return
	}

}
