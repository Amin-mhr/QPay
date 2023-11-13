package main

import (
	"fmt"
	"my-part/database"
)

func main() {
	db := database.NewGormPostgres()
	database.Migrate(db)
	fmt.Println(db)
}
