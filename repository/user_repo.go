package repository

import (
	"database/sql"
	"log"

	"github.com/rafialariq/auth-jwt/model"
)

func CheckCredential(username string, db *sql.DB) model.User {
	var user model.User

	query := "SELECT username, password FROM users WHERE username = $1;"
	row := db.QueryRow(query, username)

	if err := row.Scan(&user.Username, &user.Password); err != nil {
		log.Println(err)
	}

	if user.Username == "" {
		return user
	}

	return user
}
