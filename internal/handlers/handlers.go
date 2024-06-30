package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK,
		map[string]string{
			"message": "Investment simulator API made in Go, Postgres, Echo and more.",
			"description": `
			Start by creating a user at POST /user (name, password) add (creits) on /user/credits and create a 
			new investment by first POST /user/admin to gain admin privileges and then POST (ticker, name, minimumInvestment)
			/investment and then investing by POST /user/invest(ticker, credits) /user/invest 
			`,
		})
}
