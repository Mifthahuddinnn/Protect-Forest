package main

import (
	"forest/drivers/database"
	"forest/handler"
	"forest/repositories"
	"forest/usecases"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// User
	userRepo := repositories.UserRepository{DB: db}
	userUseCase := usecases.UserUseCase{Repo: userRepo}
	useHandler := handler.UserHandler{UserUseCase: userUseCase}

	// Admin
	adminRepo := repositories.RepositoryAdmin{DB: db}
	adminUseCase := usecases.AdminUseCase{Repo: adminRepo}
	adminHandler := handler.AdminHandler{AdminUseCase: adminUseCase}

	// Report
	//reportRepo := repositories.RepoReport{DB: db}
	//reportUseCase := usecases.ReportUseCase{Repo: &reportRepo}
	//reportHandler := handler.ReportHandler{ReportUseCase: reportUseCase}

	e := echo.New()

	// Register Login User
	e.GET("/users", useHandler.GetUsers)
	e.GET("/users/:id", useHandler.GetUserByID)
	e.POST("/users/register", useHandler.RegisterUser)
	e.POST("/users/login", useHandler.LoginUser)

	// Report User
	e.POST("/users/update", useHandler.UpdateUser)

	// Register Login Admin
	e.POST("/admin/register", adminHandler.RegisterAdmin)
	e.POST("/admin/login", adminHandler.LoginAdmin)

	e.Logger.Fatal(e.Start(":8000"))

}
