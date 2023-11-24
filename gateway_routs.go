package main

import (
	"github.com/labstack/echo/v4"
	"qpay/services"
)

func GatewayRouts(server *echo.Echo) {
	server.POST("/gateway", services.HandleCreateGateway)
	server.POST("/buy-gateway", services.HandleBuyGateway)

}
