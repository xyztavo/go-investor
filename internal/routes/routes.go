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
	// Only authenticated users can make requests to those routes
	e.GET("/user", handlers.GetUser, middlewares.GetAuth)
	e.POST("/user/credits", handlers.AddUserCredits, middlewares.GetAuth)
	e.POST("/user/investment", handlers.Invest, middlewares.GetAuth)

	// !!!! THIS ROUTE GRANTS ADMIN ACCESS TO ANY AUTHENTICATED USER
	e.POST("/user/admin", handlers.SetAdmin, middlewares.GetAuth)
	AdminRoutes(e)
}

func AdminRoutes(e *echo.Echo) {
	// Only admins can make requests to those routes
	e.POST("/investment", handlers.CreateInvestment, middlewares.AdminAuth)
	e.PATCH("/investment", handlers.UpdateInvestments, middlewares.AdminAuth)
	e.DELETE("/investment", handlers.DeleteInvestments, middlewares.AdminAuth)

	e.GET("/users", handlers.GetUsers, middlewares.AdminAuth)
	e.GET("/users/investments", handlers.GetUsersInvestments, middlewares.AdminAuth)
}
