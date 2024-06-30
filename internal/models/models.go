package models

import "github.com/golang-jwt/jwt/v5"

type JwtUserClaims struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type Investment struct {
	Id                string `json:"id"`
	Name              string `json:"name" validate:"required"`
	Ticker            string `json:"ticker" validate:"required"`
	MinimumInvestment int    `json:"minimumInvestment" validate:"required"`
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Credits  int    `json:"credits"`
}
type AuthUser struct {
	Id       string `json:"id" validate:"required"`
	Name     string `json:"name"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
	Credits  int    `json:"credits"`
}
type CreateUserStruct struct {
	Id       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role"`
	Credits  int    `json:"credits"`
}

type UserCreditsStruct struct {
	Credits int `json:"credits" validate:"required"`
}

type InvestBody struct {
	Ticker  string `json:"ticker" validate:"required"`
	Credits int    `json:"credits" validate:"required"`
}
