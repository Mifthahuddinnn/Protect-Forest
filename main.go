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
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// User
	userRepo := user2.Repository{DB: db}
	userUseCase := user.UserUseCase{Repo: userRepo}
	userHandler := handler.UserHandler{UserUseCase: userUseCase}

	// Admin
	adminRepo := admin2.Repository{DB: db}
	adminUseCase := admin.UseCaseAdmin{Repo: adminRepo}
	adminHandler := handler.AdminHandler{AdminUseCase: adminUseCase}

	// Report
	reportRepo := report.Repository{DB: db}
	reportUseCase := report2.ReportUseCase{
		Repo:     reportRepo,
		UserRepo: userRepo,
	}
	reportHandler := handler.ReportHandler{ReportUseCase: reportUseCase}

	e := echo.New()

	// Public routes
	e.GET("/users", userHandler.GetUsers)
	e.GET("/users/:id", userHandler.GetUserByID)
	e.POST("/users/register", userHandler.RegisterUser)
	e.POST("/users/login", userHandler.LoginUser)
	e.POST("/admin/register", adminHandler.RegisterAdmin)
	e.POST("/admin/login", adminHandler.LoginAdmin)

	r := e.Group("/api")
	r.Use(middleware.AuthMiddleware)
	r.POST("/reports/approve/:id", reportHandler.ApproveReport)
	r.POST("/users/redeem/:id", userHandler.RedeemPoints)
	r.POST("/reports", reportHandler.CreateReport)
	r.GET("/news", userHandler.GetNews)

	// admin
	r.GET("/reports", reportHandler.GetReports)
	r.GET("/reports/:id", reportHandler.GetReportByID)
	r.DELETE("/reports/:id", reportHandler.DeleteReport)

	e.Logger.Fatal(e.Start(":0000"))

}
