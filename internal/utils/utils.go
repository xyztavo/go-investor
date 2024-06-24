package utils

import (
	"errors"
	"strings"
	"teste/cmd/configs"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtUserClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func GetIdFromToken(c echo.Context) (userId string, err error) {
	authKey := c.Request().Header.Get("Authorization")
	authType := strings.Split(authKey, " ")[0]
	if authType != "Bearer" {
		return "", errors.New("auth type invalid, expected: Bearer {token}")
	}
	jwtFromHeader := strings.Split(authKey, " ")[1]
	// check if jwt null
	if jwtFromHeader == "" {
		return "", errors.New("not jwt token found")
	}
	var claims JwtUserClaims
	token, err := jwt.ParseWithClaims(jwtFromHeader, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.GetJwtSecret()), nil
	})
	if err != nil {
		return "", errors.New("could not parse jwt")
	}
	// check if token is valid
	if !token.Valid {
		return "", errors.New("invalid jwt")
	}
	return claims.UserId, nil
}
