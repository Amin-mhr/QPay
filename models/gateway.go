package models

import (
	"gorm.io/gorm"
)

type Gateway struct {
	gorm.Model
	UrlAddress       string
	UserID           uint
	User             User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AccountID        uint
	Account          Account `gorm:"foreignKey:AccountID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PricingID        uint
	Pricing          Pricing `gorm:"foreignKey:PricingID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TransactionCount int
}
