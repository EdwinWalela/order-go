package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name    string
	Phone   string
	Address string
}
