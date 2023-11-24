package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"qpay/services"
)

func SignUpRoutes(server *echo.Echo) {
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.POST("/signup", services.RegisterHandler(services.UserInterfaceService{}))
	server.POST("/login", services.LoginHandler(services.UserInterfaceService{}))
	server.GET("/home", services.Authentication, services.AuthMiddleware)
}
