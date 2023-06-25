package models

type Order struct {
	DeliveryFee float64
	Note        string
	Products    []Product
	Customer    Customer
}
