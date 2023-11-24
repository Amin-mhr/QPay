package services

import (
	"net/http"
	"qpay/models"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TransactionStatus string

type Message struct {
	Message string
}

func (s TransactionStatus) isValid() bool {
	_, exist := models.ValidStatuses[models.TransactionStatus(s)]
	return exist
}

//type Transaction struct {
//	ID                    uint              `gorm:"primary_key"`
//	GatewayId             uint              `gorm:"primary_key"`
//	CustomerAccountNumber int64             `gorm:"type:int"`
//	CustomerExpireDate    time.Time         `gorm:"type:date"`
//	Status                TransactionStatus `gorm:"type:string"`
//	TransactionDate       time.Time         `gorm:"type:timestamp"`
//}

type TransactionServiceInterface interface {
	List(status string) ([]*models.Transaction, error)
	FilterTransactions(date *time.Time, amount *float64) ([]*models.Transaction, error)
}

type transactionService struct {
	db *gorm.DB
}

func NewTransactionService(db *gorm.DB) TransactionServiceInterface {
	return &transactionService{db: db}
}

func (t *transactionService) List(status string) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	result := t.db.Where("status= ?", models.StatusUncompleted).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func (t *transactionService) FilterTransactions(date *time.Time, amount *float64) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	query := t.db.Model(&models.Transaction{})
	if date != nil {
		query = query.Where("transaction_date = ?", date)
	}

	if amount != nil {
		query = query.Where("amount = ?", amount)
	}

	err := query.Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func ListHandler(service TransactionServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		status := TransactionStatus(c.QueryParam("status"))
		if status.isValid() || status == "" {
			transactions, err := service.List(string(status))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			return c.JSON(http.StatusOK, transactions)
		} else {
			return c.JSON(http.StatusBadRequest, Message{Message: "The status is invalid"})
		}
	}
}

func FilterTransactionHandler(service TransactionServiceInterface) echo.HandlerFunc {
	return func(c echo.Context) error {
		var dateFilter *time.Time
		var amountFilter *float64

		dateString := c.QueryParam("date")
		if dateString != "" {
			date, dateErr := time.Parse("2006-01-02 15:04:05", dateString)
			if dateErr == nil {
				dateFilter = &date
			} else {
				return c.JSON(http.StatusBadRequest, Message{Message: "The date value is invalid"})
			}
		}

		amountString := c.QueryParam("amount")
		if amountString != "" {
			amount, amountErr := strconv.ParseFloat(amountString, 64)
			if amountErr == nil {
				amountFilter = &amount
			} else {
				return c.JSON(http.StatusBadRequest, Message{Message: "The amount is invalid"})
			}
		}

		if dateString != "" || amountString != "" {
			transactions, err := service.FilterTransactions(dateFilter, amountFilter)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
			return c.JSON(http.StatusOK, transactions)
		} else {
			return c.JSON(http.StatusBadRequest, Message{Message: "The filter is invalid"})
		}
	}
}
