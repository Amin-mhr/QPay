package routes

import (
	"my-part/controler"

	echo "github.com/labstack/echo/v4"
)

func SetUp(e* echo.Echo){
    e.POST("/api/register",controller.Register)
	e.POST("/api/login", controller.Login)
	e.GET("/api/user", controller.User)
}