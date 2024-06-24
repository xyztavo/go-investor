package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	loadErr := godotenv.Load(curDir + "/.env")
	if loadErr != nil {
		log.Fatalln("can't load env file from current directory: " + curDir)
	}
}
func GetDbConnectionString() (dbConnStr string) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("postgres://%v:%v@localhost:5432/%v?sslmode=disable", dbUser, dbPassword, dbName)
}

func GetPort() string {
	return os.Getenv("PORT")
}

func GetJwtSecret() string {
	return os.Getenv("JWT_SECRET")
}
