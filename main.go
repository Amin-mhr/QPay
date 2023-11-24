package main

import (
	"github.com/labstack/echo/v4"
	"qpay/database"
)

func main() {
	db := database.NewGormPostgres()
	database.Migrate(db)

	server := echo.New()
	GatewayRouts(server)

	server.Logger.Fatal(server.Start(":8000"))

}
