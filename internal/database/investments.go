package database

import (
	"errors"
	"teste/internal/models"
)

func InsertInvestment(Investment *models.Investment) (createdInverstmentId string, err error) {
	if err = db.QueryRow(`
	INSERT INTO investments 
	(name, ticker, minimum_investment) VALUES ($1, $2, $3)
	RETURNING id`, Investment.Name, Investment.Ticker, Investment.MinimumInvestment).Scan(&createdInverstmentId); err != nil {
		return "", errors.New("could not create investment " + err.Error())
	}
	return createdInverstmentId, nil
}

func GetInvestments() (Investments []models.Investment, err error) {
	rows, err := db.Query("SELECT * FROM investments")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var Investment models.Investment
		rows.Scan(&Investment.Id, &Investment.Name, &Investment.Ticker, &Investment.MinimumInvestment)
		Investments = append(Investments, Investment)
	}
	return Investments, nil
}
