package utils

import (
	"errors"
	"strings"
	"teste/cmd/configs"
	"teste/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtUserClaims struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GetIdFromToken(c echo.Context) (userId string, err error) {
	authKey := c.Request().Header.Get("Authorization")
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

func GetClaimsFromToken(c echo.Context) (*models.JwtUserClaims, error) {
	authKey := c.Request().Header.Get("Authorization")
	jwtFromHeader := strings.Split(authKey, " ")[1]
	// check if jwt null
	if jwtFromHeader == "" {
		return nil, errors.New("not jwt token found")
	}
	var claims models.JwtUserClaims
	token, err := jwt.ParseWithClaims(jwtFromHeader, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.GetJwtSecret()), nil
	})
	if err != nil {
		return nil, errors.New("could not parse jwt")
	}
	// check if token is valid
	if !token.Valid {
		return nil, errors.New("invalid jwt")
	}

	user := &models.JwtUserClaims{
		UserId: claims.UserId,
		Role:   claims.Role,
	}
	return user, nil
}
