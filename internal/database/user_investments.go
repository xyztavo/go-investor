package database

import (
	"errors"
	"teste/internal/models"
)

func GetUsersInvestments() (investments []models.UsersInvestments, err error) {
	rows, err := db.Query("SELECT * FROM users_investments")
	if err != nil {
		return nil, errors.New("could not get users investments")
	}
	var investment models.UsersInvestments
	for rows.Next() {
		rows.Scan(&investment.UserInvestmentId, &investment.UserId, &investment.Ticker, &investment.Amount)
		investments = append(investments, investment)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("could not scan user invesments")
	}
	rows.Close()
	return investments, nil
}
