package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	DeliveryFee float64
	Note        string
}
