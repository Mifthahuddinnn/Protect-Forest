package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

//func Connect() (*gorm.DB, error) {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatalf("Error loading .env file: %v", err)
//	}
//
//	dbHost := os.Getenv("DB_HOST")
//	dbPort := os.Getenv("DB_PORT")
//	dbUser := os.Getenv("DB_USER")
//	dbPass := os.Getenv("DB_PASS")
//	dbDatabase := os.Getenv("DB_NAME")
//
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//		dbUser, dbPass, dbHost, dbPort, dbDatabase)
//
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		return nil, err
//	}
//
//	return db, nil
//}

func ConnectDB(config Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
