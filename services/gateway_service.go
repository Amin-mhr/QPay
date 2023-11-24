package services

import (
	"github.com/labstack/echo/v4"
	"my-part/models"
)

func HandleCreateGateway(c echo.Context) error {
	var gateway models.Gateway
	err := (&echo.DefaultBinder{}).BindBody(c, &gateway)
	if err != nil {
		return err
	}

	//result, err := database.PostUser(user)
	return nil
}
