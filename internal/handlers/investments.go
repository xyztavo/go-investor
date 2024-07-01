package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"teste/internal/database"
	"teste/internal/models"
	"teste/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func validateNoSpaces(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return !strings.Contains(value, " ")
}

func CreateInvestment(c echo.Context) error {
	investment := new(models.CreateInvestment)
	if err := json.NewDecoder(c.Request().Body).Decode(&investment); err != nil {
		return echo.ErrInternalServerError
	}
	validate := validator.New()
	validate.RegisterValidation("noSpaces", validateNoSpaces)
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

func Invest(c echo.Context) error {
	// validate stuff
	id, err := utils.GetIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	body := new(models.InvestBody)
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user, err := database.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "could not find user with id "+id)
	}
	investment, err := database.GetInvestment(body.Ticker)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "can not find ticker: "+body.Ticker)
	}
	if user.Credits < investment.MinimumInvestment {
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf(`you only have %v credits but the minimum amount to invest in %v is %v`, user.Credits, investment.Ticker, investment.MinimumInvestment))
	}
	if body.Credits < investment.MinimumInvestment {
		return echo.NewHTTPError(http.StatusBadRequest,
			fmt.Sprintf(`you want to invest only %v credits but the minimum amount to invest in %v is %v`, user.Credits, investment.Ticker, investment.MinimumInvestment))
	}
	// insert and check err
	if err = database.InsertUserInvestment(user.Id, body.Ticker); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// remove credits from invested amount and check err
	if err = database.RemoveUserInvestedCredits(body.Credits, user.Id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "invested with ease!"})
}

func GetUsersInvestments(c echo.Context) error {
	investments, err := database.GetUsersInvestments()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, investments)
}
