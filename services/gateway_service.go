package services

import (
	"github.com/labstack/echo/v4"
	"my-part/database"
	"my-part/models"
	"net/http"
)

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
