package main

import (
	"forest/drivers/database"
	"forest/handler"
	"forest/repositories"
	"forest/usecases"
	"forest/utils"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	// Load Env
	utils.LoadEnv()

	// Init DB
	config := utils.InitConfigMysql()
	db, err := database.ConnectDB(config)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// User
	userRepo := repositories.UserRepository{DB: db}
	userUseCase := usecases.UserUseCase{Repo: userRepo}
	useHandler := handler.UserHandler{UserUseCase: userUseCase}

	// Admin
	adminRepo := repositories.RepositoryAdmin{DB: db}
	adminUseCase := usecases.AdminUseCase{Repo: adminRepo}
	adminHandler := handler.AdminHandler{AdminUseCase: adminUseCase}

	e := echo.New()

	// Register Login User
	e.GET("/users", useHandler.GetUsers)
	e.GET("/users/:id", useHandler.GetUserByID)
	e.POST("/users/register", useHandler.RegisterUser)
	e.POST("/users/login", useHandler.LoginUser)

	// Register Login Admin
	e.POST("/admin/register", adminHandler.RegisterAdmin)
	e.POST("/admin/login", adminHandler.LoginAdmin)

	e.Logger.Fatal(e.Start(":8000"))

}
