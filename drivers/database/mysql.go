package database

import (
	"fmt"
	"forest/entities"
	"log"
	"os"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect() (*gorm.DB, error) {
	envProject := os.Getenv("ENV_PROJECT")
	envProject = strings.ReplaceAll(envProject, "\"", "")
	envs := strings.Split(envProject, " ")

	var envMap map[string]string
	envMap = make(map[string]string)
	for _, env := range envs {
		keyValue := strings.Split(env, "=")
		envMap[keyValue[0]] = keyValue[1]
	}

	dbHost := envMap["DB_HOST"]
	dbPort := envMap["DB_PORT"]
	dbUser := envMap["DB_USER"]
	dbPass := envMap["DB_PASS"]
	dbDatabase := envMap["DB_NAME"]

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	if err := db.AutoMigrate(&entities.User{}, &entities.Admin{}, &entities.Report{}, &entities.Redeem{}, &entities.Balance{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
		return nil, err
	}

	return db, nil
}
