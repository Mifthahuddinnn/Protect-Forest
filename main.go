package main

import (
	"forest/drivers/database"
	"forest/handler"
	"forest/middleware"
	admin2 "forest/repositories/admin"
	"forest/repositories/report"
	user2 "forest/repositories/user"
	"forest/usecases/admin"
	report2 "forest/usecases/report"
	"forest/usecases/user"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// User
	userRepo := user2.Repository{DB: db}
	userUseCase := user.UserUseCase{Repo: userRepo}
	useHandler := handler.UserHandler{UserUseCase: userUseCase}

	// Admin
	adminRepo := admin2.Repository{DB: db}
	adminUseCase := admin.UseCaseAdmin{Repo: adminRepo}
	adminHandler := handler.AdminHandler{AdminUseCase: adminUseCase}

	// Report
	reportRepo := report.Repository{DB: db}
	reportUseCase := report2.ReportUseCase{Repo: reportRepo}
	reportHandler := handler.ReportHandler{ReportUseCase: reportUseCase}

	e := echo.New()

	// Register Login User
	e.GET("/users", useHandler.GetUsers)
	e.GET("/users/:id", useHandler.GetUserByID)
	e.POST("/users/register", useHandler.RegisterUser)
	e.POST("/users/login", useHandler.LoginUser)

	r := e.Group("")
	r.Use(middleware.JWTMiddleware)

	// Reporting
	r.POST("/reports", reportHandler.CreateReport)

	// Register Login Admin
	e.POST("/admin/register", adminHandler.RegisterAdmin)
	e.POST("/admin/login", adminHandler.LoginAdmin)

	e.Logger.Fatal(e.Start(":8000"))

}
