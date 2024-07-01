package models

import "github.com/golang-jwt/jwt/v5"

type JwtUserClaims struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type Investment struct {
	Name              string `json:"name"`
	Ticker            string `json:"ticker"`
	MinimumInvestment int    `json:"minimumInvestment"`
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Credits  int    `json:"credits"`
}

type UsersInvestments struct {
	UserInvestmentId string `json:"userInvestmentId"`
	UserId           string `json:"userId"`
	Ticker           string `json:"ticker"`
}
