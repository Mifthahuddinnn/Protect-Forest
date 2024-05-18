package database

import (
	"fmt"
	"forest/entities"
	"github.com/joho/godotenv"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	// Attempt to load .env file, but don't fail if it doesn't exist
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Get environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbDatabase := os.Getenv("DB_NAME")

	// Check if any of the environment variables are empty
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbDatabase == "" {
		log.Fatalf("One or more environment variables are not set: DB_HOST=%s, DB_PORT=%s, DB_USER=%s, DB_PASS=%s, DB_NAME=%s", dbHost, dbPort, dbUser, dbPass, dbDatabase)
	}

	// Create DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbDatabase)

	// Open connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	// AutoMigrate database entities
	if err := db.AutoMigrate(&entities.User{}, &entities.Admin{}, &entities.Report{}, &entities.Redeem{}, &entities.Balance{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
		return nil, err
	}

	return db, nil
}
