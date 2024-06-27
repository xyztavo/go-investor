package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK,
		map[string]string{"message": "this is a simple crud project."})
}
