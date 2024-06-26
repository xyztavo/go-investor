package models

import "github.com/golang-jwt/jwt/v5"

type JwtUserClaims struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type Investments struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Ticker string `json:"ticker"`
}
