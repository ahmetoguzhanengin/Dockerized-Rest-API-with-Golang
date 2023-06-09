package config

import (
	"fmt"
	"time"

	"github.com/ahmetoguzhanengin/RestAPI/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://depixen:depixen-pass@postgresql_db/postgres?sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	start := time.Now()
	for sqlDB.Ping() != nil {
		if start.After(start.Add(10 * time.Second)) {
			fmt.Println("Failed to connect DB after 10 seconds")
			break
		}
	}
	fmt.Println("Connected to Db: ", sqlDB.Ping() == nil)
	db.AutoMigrate(&models.Card{})
}
