package handlers

import (
	"encoding/json"
	"net/http"
	"teste/cmd/configs"
	"teste/internal/database"
	"teste/internal/models"
	"teste/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func HelloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK,
		map[string]string{"message": "this is a simple crud project."})
}

func CreateUser(c echo.Context) error {
	// create var user
	user := new(models.CreateUserStruct)
	// decode the req body to the user var
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "an error has occurred")
	}
	validate := validator.New()
	// validate the struct
	if err := validate.Struct(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "an error has occurred")
	}
	user.Password = string(hashedPassword)
	userId, err := database.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "an error has occured while creating the user in the database")
	}
	jwtKey := []byte(configs.GetJwtSecret())
	// create jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JwtUserClaims{
		UserId: userId,
		Role:   "investor",
	})
	// sign jwt secret
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "an error has occurred")
	}
	return c.JSON(http.StatusCreated, map[string]string{"token": signedToken})
}

func AuthUser(c echo.Context) error {
	user := new(models.User)
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error while decoding json body")
	}
	validate := validator.New()
	// validate the struct
	if err := validate.Struct(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// get user from db
	userFromDb, err := database.GetUserById(user.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	// compare hash password with password from db
	err = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(user.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "password does not match")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.JwtUserClaims{
		UserId: user.Id,
		Role:   user.Role,
	})
	jwtKey := []byte(configs.GetJwtSecret())
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "error while signing jwt")
	}
	return c.JSON(http.StatusOK, map[string]string{"token": signedToken})
}

func GetUsers(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	idFromToken, err := utils.GetIdFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	user, err := database.GetUserById(idFromToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user by id not found")
	}
	return c.JSON(http.StatusOK, map[string]string{"id": user.Id, "name": user.Name, "role": user.Role})
}

func CreateInvestment(c echo.Context) error {
	investment := new(models.Investments)
	err := json.NewDecoder(c.Request().Body).Decode(&investment)
	if err != nil {
		return echo.ErrInternalServerError
	}
	user, err := utils.GetClaimsFromToken(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	if user.Role != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "must be admin to view this")
	}
	createdInverstmentId, err := database.InsertInvestment(investment.Name, investment.Ticker)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]string{"createdInvestmentId": createdInverstmentId})
}
