package main

import (
	"github.com/labstack/echo/v4"
	"my-part/database"
	"my-part/services"
)

func main() {
	db := database.NewGormPostgres()
	database.Migrate(db)

	server := echo.New()
	server.POST("/gateway", services.HandleCreateGateway)

	server.Logger.Fatal(server.Start(":8000"))

}
