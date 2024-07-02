package database

import (
	"errors"
	"teste/internal/models"
)

func InsertInvestment(Investment *models.CreateInvestment) (createdInvestmentTicker string, err error) {
	if err = db.QueryRow(`
	INSERT INTO investments 
	(name, ticker, minimum_investment) VALUES ($1, $2, $3)
	RETURNING ticker`, Investment.Name, Investment.Ticker, Investment.MinimumInvestment).Scan(&createdInvestmentTicker); err != nil {
		return "", errors.New("could not create investment " + err.Error())
	}
	return createdInvestmentTicker, nil
}

func GetInvestments() (Investments []models.Investment, err error) {
	rows, err := db.Query("SELECT * FROM investments")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var Investment models.Investment
		rows.Scan(&Investment.Ticker, &Investment.Name, &Investment.MinimumInvestment)
		Investments = append(Investments, Investment)
	}
	return Investments, nil
}

func GetInvestment(ticker string) (Investment models.Investment, err error) {
	err = db.QueryRow("SELECT * FROM investments WHERE ticker = $1", ticker).
		Scan(&Investment.Ticker, &Investment.Name, &Investment.MinimumInvestment)
	if err != nil {
		return Investment, err
	}
	return Investment, err
}

func InsertUserInvestment(userId string, ticker string, amount int) error {
	res, err := db.Exec(`INSERT INTO users_investments (user_id, ticker, amount) VALUES ($1, $2, $3)`, userId, ticker, amount)
	if err != nil {
		return errors.New("could not insert user investment")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("cannot see rowsaffected for some reason")
	}
	if rowsAffected < 1 {
		return errors.New("no rows were affected")
	}
	return nil
}

func RemoveInvestment(ticker string) error {
	res, err := db.Exec("DELETE FROM investments WHERE ticker = $1", ticker)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return errors.New("could not delete investment")
	}
	return nil
}

func UpdateInvestment(updateInvestmentBody *models.UpdateInvestment) error {
	res, err := db.Exec("UPDATE investments SET name = $1, minimum_investment = $2 WHERE ticker = $3", updateInvestmentBody.Name, updateInvestmentBody.MinimumInvestment, updateInvestmentBody.Ticker)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return errors.New("could not update investment")
	}
	return nil
}

func RemoveUserInvestedCredits(amount int, userId string) error {
	res, err := db.Exec("UPDATE users SET credits = credits - $1 WHERE id = $2", amount, userId)
	if err != nil {
		return errors.New("could not remove invested credits from user")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.New("error while trying to get rowsaffected from removing user invested credits")
	}
	if rowsAffected < 1 {
		return errors.New("no rows were affected while trying to remove user invested credits")
	}
	return nil
}
