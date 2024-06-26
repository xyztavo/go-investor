package database

import (
	"database/sql"
	"log"
	"teste/cmd/configs"

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
		role VARCHAR(40) NOT NULL
		);
		CREATE TABLE IF NOT EXISTS investments (
		id SERIAL PRIMARY KEY,
		name VARCHAR(40) NOT NULL,
		ticker VARCHAR(40) NOT NULL
		);
	`)
	if err != nil {
		return err
	}
	return nil
}
