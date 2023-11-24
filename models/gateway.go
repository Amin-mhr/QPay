package models

import (
	"gorm.io/gorm"
	"time"
)

type Gateway struct {
	gorm.Model
	Blocked          bool
	AlwaysBlocked    bool
	BlockTime        time.Time
	UnblockTime      time.Time
	UrlAddress       string `gorm:"unique"`
	UserID           uint
	User             User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	AccountID        uint
	Account          Account `gorm:"foreignKey:AccountID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PricingID        uint
	Pricing          Pricing `gorm:"foreignKey:PricingID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TransactionCount int
}
