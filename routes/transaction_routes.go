package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"qpay/services"
)

func TransactionRoutes(server *echo.Echo, db *gorm.DB) {
	transactionService := services.NewTransactionService(db)
	server.GET("/transaction/list", services.ListHandler(transactionService))
	server.GET("/transaction/filter", services.FilterTransactionHandler(transactionService))
}
