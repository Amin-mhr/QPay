package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	IsCompany bool
	Name      string
	Email     string
	Password  string
}
