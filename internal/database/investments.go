package database

import (
	"errors"
	"teste/internal/models"
)

func InsertInvestment(name string, ticker string) (createdInverstmentId string, err error) {
	if err = db.QueryRow(`
	INSERT INTO investments 
	(name, ticker) VALUES ($1, $2)
	RETURNING id`, name, ticker).Scan(&createdInverstmentId); err != nil {
		return "", errors.New("could not create investment " + err.Error())
	}
	return createdInverstmentId, nil
}

func GetInvestments() (Investments []models.Investments, err error) {
	rows, err := db.Query("SELECT * FROM investments")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var Investment models.Investments
		rows.Scan(&Investment.Id, &Investment.Name, &Investment.Ticker)
		Investments = append(Investments, Investment)
	}
	return Investments, nil
}
