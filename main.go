package main

import (
	"github.com/labstack/echo/v4"
	"qpay/database"
	"qpay/routes"
)

func main() {
	db := database.NewGormPostgres()
	database.Migrate(db)

	server := echo.New()
	routes.GatewayRouts(server)
	routes.TransactionRoutes(server, db)
	routes.SignUpRoutes(server)
	routes.AdminRoutes(server)

	server.Logger.Fatal(server.Start(":8000"))

}
