package database

import (
	"my-part/models"
	"net/http"
	"time"
)

func PostGateway(gateway models.Gateway) (int, error) {
	db := NewGormPostgres()
	gateway.CreatedAt = time.Now()
	result := db.Create(&gateway)

	if result.Error != nil {
		return http.StatusBadRequest, result.Error
	}
	return http.StatusOK, nil
}
