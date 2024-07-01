package database

import (
	"database/sql"
	"log"
	"teste/configs"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	// weird but works
	var err error
	db, err = sql.Open("postgres", configs.GetDbConnectionString())
	if err != nil {
		log.Fatal(err)
	}
}

func NewConn() *sql.DB {
	// this reuses the connection
	return db
}

func Migrate() error {
	_, err := db.Exec(`
	 	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(40) NOT NULL,
		password VARCHAR(200) NOT NULL,
		role VARCHAR(40) NOT NULL,
		credits INT 
		);
		CREATE TABLE IF NOT EXISTS investments (
		ticker VARCHAR(40) PRIMARY KEY,
		name VARCHAR(40) NOT NULL,
		minimum_investment INT NOT NULL
		);
		CREATE TABLE IF NOT EXISTS users_investments (
		user_invesment_id SERIAL PRIMARY KEY, 
		user_id INT REFERENCES users(id) ON DELETE CASCADE,
		ticker VARCHAR(40) REFERENCES investments(ticker) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return err
	}
	return nil
}
