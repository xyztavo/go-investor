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
	e.POST("/auth", handlers.AuthUser)

	AuthRequiredRoutes(e)
}

func AuthRequiredRoutes(e *echo.Echo) {
	// Any user with a valid jwt can access those routes
	e.GET("/user", handlers.GetUser, middlewares.GetAuth)
	e.POST("/user/credits", handlers.AddUserCredits, middlewares.GetAuth)
	e.POST("/user/investment", handlers.Invest, middlewares.GetAuth)

	// !DO NOT LET THAT GO IN PROD! this routes grants admin access to any auth user,
	e.POST("/user/admin", handlers.SetAdmin, middlewares.GetAuth)
	AdminRoutes(e)
}

func AdminRoutes(e *echo.Echo) {
	// Only admins can get this routes
	e.POST("/investment", handlers.CreateInvestment, middlewares.AdminAuth)
	e.GET("/users", handlers.GetUsers, middlewares.AdminAuth)
	e.GET("/users/investments", handlers.GetUsersInvestments, middlewares.AdminAuth)
}
