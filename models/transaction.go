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
	Status                TransactionStatus
	Amount                float64
	TransactionDate       time.Time
}

type TransactionStatus string

const (
	StatusUncompleted        TransactionStatus = "Uncompleted"
	StatusUnsuccessful       TransactionStatus = "unsuccessful"
	StatusFailed             TransactionStatus = "failed"
	StatusBlocked            TransactionStatus = "blocked"
	StatusRefundToPayer      TransactionStatus = "refund"
	StatusSystemRefund       TransactionStatus = "systemRefund"
	StatusCanceled           TransactionStatus = "canceled"
	StatusRedirected         TransactionStatus = "redirected"
	StatusPending            TransactionStatus = "pending"
	StatusConfirmed          TransactionStatus = "confirmed"
	StatusDepositedRecipient TransactionStatus = "depositedRecipient"
	StatusAlreadyConfirmed   TransactionStatus = "alreadyConfirmed"
)

var ValidStatuses = map[TransactionStatus]bool{
	StatusUncompleted:        true,
	StatusUnsuccessful:       true,
	StatusFailed:             true,
	StatusBlocked:            true,
	StatusRefundToPayer:      true,
	StatusSystemRefund:       true,
	StatusCanceled:           true,
	StatusRedirected:         true,
	StatusPending:            true,
	StatusConfirmed:          true,
	StatusDepositedRecipient: true,
	StatusAlreadyConfirmed:   true,
}
