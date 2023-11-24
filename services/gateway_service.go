package services

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"qpay/database"
	"qpay/models"
	"strconv"
)

// For gateway which are not specific add name to last name and put it as gateway

func HandleCreateGateway(c echo.Context) error {
	msg := make(map[string]string)
	var gateway models.Gateway
	err := (&echo.DefaultBinder{}).BindBody(c, &gateway)
	if err != nil {
		return err
	}

	result, err1 := database.PostGateway(gateway)
	if err1 != nil {
		return err1
	}
	if result != http.StatusOK {
		msg["message"] = "problem inserting gateway to database"
		return c.JSON(result, msg)
	}

	msg["message"] = "gateway added to database successfully"
	return c.JSON(http.StatusOK, msg)
}

// HandleBuyGateway TODO:
// handle money to minus from account
func HandleBuyGateway(c echo.Context) error {
	msg := make(map[string]string)
	var gateway models.Gateway
	err := (&echo.DefaultBinder{}).BindBody(c, &gateway)
	if err != nil {
		return err
	}
	result, err1 := database.PostCustomUrlGateway(gateway)
	if err1 != nil {
		return err1
	}
	if result != http.StatusOK {
		msg["message"] = "problem inserting gateway to database"
		return c.JSON(http.StatusBadRequest, msg)
	}
	return nil
}

func HandleChangeAccount(c echo.Context) error {
	msg := make(map[string]string)
	_ = msg
	gatewayId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	gatewayAccount := c.Param("account")

	result, err := database.UpdateGatewayAccount(gatewayId, gatewayAccount)
	if result != http.StatusOK {
		msg["message"] = "problem changing gateway account number"
		return c.JSON(http.StatusBadRequest, msg)
	}
	msg["message"] = "gateway account changed"
	return c.JSON(http.StatusOK, msg)
}
