package main

import (
	"fmt"
	"forest/handler"
	"forest/usecases"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbUser     = os.Getenv("DB_USER")
	dbPass     = os.Getenv("DB_PASS")
	dbDatabase = os.Getenv("DB_NAME")
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbDatabase)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// Initialize user repository and usecase
	userRepository := NewUserRepository(gormDB)
	userUseCase := usecases.UserUseCase{UserRepository: userRepository}

	// Create echo instance
	e := echo.New()

	// Use middleware for JWT authorization (replace with your own middleware logic)
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
	}))

	// Create user handler
	userHandler := handler.UserHandler{UserUseCase: userUseCase}

	// Define API endpoints
	e.GET("/users", userHandler.GetUsers)
	e.GET("/users/:id", userHandler.GetUserByID)
	e.POST("/users", userHandler.RegisterUser)
	e.POST("/login", userHandler.LoginUser)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

func NewUserRepository(db *gorm.DB) usecases.UserRepository {
	// Implement user repository logic here
	return nil
}
