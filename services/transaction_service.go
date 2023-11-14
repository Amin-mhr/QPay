package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

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

type Transaction struct {
	ID                    uint              `gorm:"primary_key"`
	GatewayId             uint              `gorm:"primary_key"`
	CustomerAccountNumber int64             `gorm:"type:int"`
	CustomerExpireDate    time.Time         `gorm:"type:date"`
	Status                TransactionStatus `gorm:"type:string"`
	TransactionDate       time.Time         `gorm:"type:timestamp"`
}

type TransactionServiceInterface interface {
	UncompletedList() ([]*Transaction, error)
	UnsuccessfulList() ([]*Transaction, error)
	FailedList() ([]*Transaction, error)
	BlockedList() ([]*Transaction, error)
	RefundToPayerList() ([]*Transaction, error)
	//SystemRefundList() ([]*Transaction, error)
	//CanceledList() ([]*Transaction, error)
	//RedirectedList() ([]*Transaction, error)
	//PendingList() ([]*Transaction, error)
	//ConfirmedList() ([]*Transaction, error)
	//DepositedToRecipientList() ([]*Transaction, error)
	//AlreadyConfirmedList() ([]*Transaction, error)
	//List() ([]*Transaction, error)
}

type transactionService struct {
	db *gorm.DB
}

func NewTransactionService(db *gorm.DB) TransactionServiceInterface {
	return &transactionService{db: db}
}

func (t *transactionService) UncompletedList() ([]*Transaction, error) {
	var transactions []*Transaction
	result := t.db.Where("status= ?", StatusUncompleted).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func (t *transactionService) UnsuccessfulList() ([]*Transaction, error) {
	var transactions []*Transaction
	result := t.db.Where("status= ?", StatusUnsuccessful).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func (t *transactionService) FailedList() ([]*Transaction, error) {
	var transactions []*Transaction
	result := t.db.Where("status= ?", StatusFailed).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func (t *transactionService) BlockedList() ([]*Transaction, error) {
	var transactions []*Transaction
	result := t.db.Where("status= ?", StatusBlocked).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func (t *transactionService) RefundToPayerList() ([]*Transaction, error) {
	var transactions []*Transaction
	result := t.db.Where("status= ?", StatusRefundToPayer).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func uncompletedListHandler(service TransactionServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		transactions, err := service.UncompletedList()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, transactions)
	}
}

func unsuccessfulListHandler(service TransactionServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		transactions, err := service.UnsuccessfulList()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, transactions)
	}
}

func failedListHandler(service TransactionServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		transactions, err := service.FailedList()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, transactions)
	}
}

func blockedListHandler(service TransactionServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		transactions, err := service.BlockedList()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, transactions)
	}
}

func refundToPayerListHandler(service TransactionServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		transactions, err := service.RefundToPayerList()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, transactions)
	}
}

func TransactionRoutes(server *echo.Echo, db *gorm.DB) {
	transactionService := NewTransactionService(db)
	server.Get("/transaction/uncompleted", uncompletedListHandler(transactionService))
	server.Get("/transaction/unsuccessful", unsuccessfulListHandler(transactionService))
	server.Get("/transaction/failed", failedListHandler(transactionService))
	server.Get("/transaction/blocked", blockedListHandler(transactionService))
	server.Get("/transaction/refund-to-payer", refundToPayerListHandler(transactionService))
}
