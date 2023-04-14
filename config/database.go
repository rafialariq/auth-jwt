package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/rafialariq/auth-jwt/utils"
)

func ConnectDb() *sql.DB {

	dbHost := utils.DotEnv("DB_HOST")
	dbPort := utils.DotEnv("DB_PORT")
	dbUser := utils.DotEnv("DB_USER")
	dbPassword := utils.DotEnv("DB_PASS")
	dbName := utils.DotEnv("DB_NAME")
	sslMode := utils.DotEnv("SSL_MODE")

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPassword, dbName, sslMode)
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
