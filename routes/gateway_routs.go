package routes

import (
	"github.com/labstack/echo/v4"
	"qpay/services"
)

func GatewayRouts(server *echo.Echo) {
	server.POST("/gateway", services.HandleCreateGateway)
	server.POST("/buy-gateway", services.HandleBuyGateway)
	server.PATCH("/change-account/:id/:account", services.HandleChangeAccount)

}
