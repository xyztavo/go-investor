package routes

import (
	"teste/internal/handlers"
	"teste/internal/middlewares"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", handlers.HelloWorld)
	e.GET("/investments", handlers.GetInvestments)
	e.POST("/user", handlers.CreateUser)
	e.GET("/users", handlers.GetUsers)
	e.POST("/auth", handlers.AuthUser)

	AuthRequiredRoutes(e)
}

func AuthRequiredRoutes(r *echo.Echo) {
	// i want this route to use the middlewares.GetAuth
	r.GET("/user", handlers.GetUser, middlewares.GetAuth)
	r.POST("/user/credits", handlers.AddUserCredits, middlewares.GetAuth)
	r.POST("/user/admin", handlers.SetAdmin, middlewares.GetAuth)
	r.POST("/investment", handlers.CreateInvestment, middlewares.AdminAuth)
}
