package routes

import (
	"teste/internal/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", handlers.HelloWorld)
	e.POST("/user", handlers.CreateUser)
	e.GET("/users", handlers.GetUsers)
	e.GET("/user", handlers.GetUser)
	e.POST("/auth", handlers.AuthUser)
}
