package models

import "github.com/golang-jwt/jwt/v5"

type JwtUserClaims struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type Investments struct {
	Id     string `json:"id"`
	Name   string `json:"name" validate:"required"`
	Ticker string `json:"ticker" validate:"required"`
}

type User struct {
	Id       string `json:"id" validate:"required"`
	Name     string `json:"name"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}
type CreateUserStruct struct {
	Id       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
}
