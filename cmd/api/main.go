package main

import (
	"fmt"
	"log"
	"teste/configs"
	"teste/internal/database"
	"teste/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	routes.SetupRoutes(e)
	if err := database.Migrate(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("listening on http://localhost%v\n", configs.GetPort())
	e.Logger.Fatal(e.Start(configs.GetPort()))
}
