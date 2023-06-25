package db

import (
	"fmt"
	"log"

	"github.com/edwinwalela/go-order/orders/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Up(config config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Db.Host, config.Db.User, config.Db.Password, config.Db.Name, config.Db.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err.Error())
	}
	log.Println("connection to database successful")
	return db, nil
}
