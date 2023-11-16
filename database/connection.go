package database

import (
	"log"
	"net/url"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB
var once sync.Once

func NewGormPostgres() *gorm.DB {
	once.Do(func() {
		tehranTimezone, _ := time.LoadLocation("Asia/Tehran")

		// Connection configuration
		dsn := &url.URL{
			Scheme:   "postgres",
			User:     url.UserPassword("pg", "pass"),
			Host:     "localhost",
			Path:     "qpay",
			RawQuery: "sslmode=disable&timezone=" + tehranTimezone.String(),
		}

		// Convert URL to connection string
		connStr := dsn.String()

		db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to the database:", err)
		}

		instance = db
	})

	return instance
}
