package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"qpay/services"
)

func AdminRoutes(server *echo.Echo) {
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.POST("/admin/login", services.LoginAdminHandler(services.AdminInterfaceService{}))
	server.POST("/admin/blockOneGateway", services.BlockOneGateWayHandler(services.AdminInterfaceService{}))
	server.POST("/admin/blockAllGateways", services.BlockAllGateWayHandler(services.AdminInterfaceService{}))
	server.GET("/admin", services.Authentication, services.AuthMiddleware)
}
