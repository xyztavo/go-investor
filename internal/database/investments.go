package database

import "errors"

func InsertInvestment(name string, ticker string) (createdInverstmentId string, err error) {
	if err = db.QueryRow(`
	INSERT INTO TABLE investments 
	(name, ticker) VALUES ($1, $2)
	RETURNING id`, name, ticker).Scan(&createdInverstmentId); err != nil {
		return "", errors.New("could not create investment")
	}
	return createdInverstmentId, nil
}
