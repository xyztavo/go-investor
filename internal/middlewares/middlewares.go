package middlewares

import (
	"errors"
	"strings"
	"teste/cmd/configs"
	"teste/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authKey := c.Request().Header.Get("Authorization")
		authType := strings.Split(authKey, " ")[0]
		if authType != "Bearer" {
			return errors.New("auth type invalid, expected: Bearer {token}")
		}
		jwtFromHeader := strings.Split(authKey, " ")[1]
		// check if jwt null
		if jwtFromHeader == "" {
			return errors.New("not jwt token found")
		}
		var claims models.JwtUserClaims
		token, err := jwt.ParseWithClaims(jwtFromHeader, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.GetJwtSecret()), nil
		})
		if err != nil {
			return errors.New("could not parse jwt")
		}
		// check if token is valid
		if !token.Valid {
			return errors.New("invalid jwt")
		}
		return next(c)
	}
}
