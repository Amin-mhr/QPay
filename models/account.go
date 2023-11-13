package models

import (
	"gorm.io/gorm"
	"time"
)

type Account struct {
	gorm.Model
	Bank          string
	AccountNumber string
	ExpireDate    time.Time
}
