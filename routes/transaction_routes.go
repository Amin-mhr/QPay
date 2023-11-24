package routes

import (
	"qpay/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TransactionRoutes(server *echo.Echo, db *gorm.DB) {
	transactionService := services.NewTransactionService(db)
	server.GET("/transaction/list", services.ListHandler(transactionService))
	server.GET("/transaction/filter", services.FilterTransactionHandler(transactionService))
	server.POST("/transaction", services.HandleCreateTransaction(transactionService))
}
