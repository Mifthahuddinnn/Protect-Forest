package database

import (
	"fmt"
	"forest/entities"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(); err != nil {
			log.Printf("Error loading .env file: %v", err)
		} else {
			log.Println(".env file loaded successfully")
		}
	} else {
		log.Println(".env file not found, skipping loading")
	}

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
