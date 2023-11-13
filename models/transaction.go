package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	GatewayID             uint
	Gateway               Gateway `gorm:"foreignKey:GatewayID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CustomerAccountNumber string
	CustomerExpireDate    time.Time
	Status                string
	Amount                float64
	TransactionDate       time.Time
}
