package models

import "gorm.io/gorm"

type Pricing struct {
	gorm.Model
	Months           int
	TransactionLimit int
}
