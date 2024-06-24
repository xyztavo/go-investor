package handlers

import (
	"encoding/json"
	"net/http"
	"teste/cmd/configs"
	"teste/internal/database"
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

type JwtUserClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func CreateUser(c echo.Context) error {
	// create var user
	user := new(database.CreateUserStruct)
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtUserClaims{
		UserId: userId,
	})
	// sign jwt secret
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "an error has occurred")
	}
	return c.JSON(http.StatusCreated, map[string]string{"token": signedToken})
}

func AuthUser(c echo.Context) error {
	user := new(database.User)
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "an error has ocurred")
	}
	userFromDb, err := database.GetUserById(user.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userFromDb.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "password does not match")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JwtUserClaims{
		UserId: user.Id,
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
	return c.JSON(http.StatusOK, map[string]string{"id": user.Id, "name": user.Name})

}
