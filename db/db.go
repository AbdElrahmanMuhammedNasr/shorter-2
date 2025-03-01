package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"shorter/model"
	"time"
)

var DB *gorm.DB

func IntiDB() {

	dsn := "host=localhost user=postgres password=root dbname=short port=6666 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	// Set connection pool settings
	sqlDB.SetMaxOpenConns(50)                 // Maximum open connections
	sqlDB.SetMaxIdleConns(25)                 // Maximum idle connections
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Maximum connection lifetime
	sqlDB.SetConnMaxIdleTime(2 * time.Minute) // Maximum idle time for a connection

	DB.AutoMigrate(&model.Urls{})
	DB.AutoMigrate(&model.Users{})
}
