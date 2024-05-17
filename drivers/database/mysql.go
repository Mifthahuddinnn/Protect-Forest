package database

import (
	"fmt"
	"forest/entities"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {

	// Get environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbDatabase := os.Getenv("DB_NAME")

	// Create DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbDatabase)

	// Open database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Check if the database exists
	var dbName string
	err = db.Raw("SELECT DATABASE()").Scan(&dbName).Error
	if err != nil || dbName != dbDatabase {
		log.Fatalf("Database does not exist or connection failed: %v", err)
		return nil, err
	}

	// Auto-migrate the models
	if err := db.AutoMigrate(&entities.User{}, &entities.Admin{}, &entities.Report{}, &entities.Redeem{}, &entities.Balance{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
		return nil, err
	}

	return db, nil
}
