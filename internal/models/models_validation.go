package models

type CreateInvestment struct {
	Name              string `json:"name" validate:"required"`
	Ticker            string `json:"ticker" validate:"required,noSpaces,uppercase"`
	MinimumInvestment int    `json:"minimumInvestment" validate:"required"`
}

type AuthUser struct {
	Id       string `json:"id" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type CreateUser struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserCredits struct {
	Credits int `json:"credits" validate:"required"`
}

type InvestBody struct {
	Ticker  string `json:"ticker" validate:"required"`
	Credits int    `json:"credits" validate:"required"`
}
