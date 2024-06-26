package routes

import (
	"teste/internal/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", handlers.HelloWorld)
	e.POST("/user", handlers.CreateUser)
	e.GET("/users", handlers.GetUsers)
	e.POST("/auth", handlers.AuthUser)
}

func AuthRequiredRoutes(r *echo.Echo) {
	// // if i use this middleware it applies to every route, maybe read the docs later,
	// r.Use(middlewares.GetAuth)
	// r.GET("/user", handlers.GetUser)
}
