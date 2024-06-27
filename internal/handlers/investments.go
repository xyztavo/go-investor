package handlers

import (
	"encoding/json"
	"net/http"
	"teste/internal/database"
	"teste/internal/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CreateInvestment(c echo.Context) error {
	investment := new(models.Investment)
	if err := json.NewDecoder(c.Request().Body).Decode(&investment); err != nil {
		return echo.ErrInternalServerError
	}
	validate := validator.New()
	// validate the struct
	if err := validate.Struct(investment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	createdInverstmentId, err := database.InsertInvestment(investment)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]string{"createdInvestmentId": createdInverstmentId})
}

func GetInvestments(c echo.Context) error {
	investments, err := database.GetInvestments()
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, investments)
}
