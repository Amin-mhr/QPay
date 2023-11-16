package main

import (
	"fmt"
	"my-part/database"
	"my-part/routes"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := database.NewGormPostgres()
	database.Migrate(db)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
	}))

	routes.SetUp(e)
	e.Logger.Fatal(e.Start(":8000"))

	fmt.Println(db)
}
