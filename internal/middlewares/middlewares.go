package middlewares

import (
	"net/http"
	"strings"
	"teste/configs"
	"teste/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authKey := c.Request().Header.Get("Authorization")
		if authKey == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "auth token found")
		}
		jwtFromHeader := strings.Split(authKey, " ")[1]
		// check if jwt null
		if jwtFromHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "not jwt token found")
		}
		var claims models.JwtUserClaims
		token, err := jwt.ParseWithClaims(jwtFromHeader, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.GetJwtSecret()), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "could not parse jwt")
		}
		// check if token is valid
		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid jwt")
		}

		return next(c)
	}
}

func AdminAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authKey := c.Request().Header.Get("Authorization")
		if authKey == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "auth token found")
		}
		jwtFromHeader := strings.Split(authKey, " ")[1]
		// check if jwt null
		if jwtFromHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "not jwt token found")
		}
		var claims models.JwtUserClaims
		token, err := jwt.ParseWithClaims(jwtFromHeader, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.GetJwtSecret()), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "could not parse jwt")
		}
		// check if token is valid
		if !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid jwt")
		}
		user := &models.JwtUserClaims{
			Role: claims.Role,
		}
		if user.Role != "admin" {
			return echo.NewHTTPError(http.StatusUnauthorized, "must be admin")
		}
		return next(c)
	}
}
