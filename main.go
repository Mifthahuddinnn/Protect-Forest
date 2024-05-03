package main

import (
	"fmt"
	"forest/handler"
	"forest/repositories/user"
	"forest/usecases"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbDatabase := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepo := user.Repository{DB: db}
	userUseCase := usecases.UserUseCase{Repo: userRepo}
	useHandler := handler.UserHandler{UserUseCase: userUseCase}

	e := echo.New()
	e.GET("/users", useHandler.GetUsers)
	e.GET("/users/:id", useHandler.GetUserByID)
	e.POST("/users/register", useHandler.RegisterUser)
	e.POST("/users/login", useHandler.LoginUser)
	e.Logger.Fatal(e.Start(":8000"))

}
