package repository

import (
	"database/sql"
	"log"

	"github.com/rafialariq/auth-jwt/model"
)

// type UserRepo interface {
// 	CheckCredential(username string) any
// }

// type userRepo struct {
// 	db *sql.DB
// }

func CheckCredential(username string, db *sql.DB) any {
	var user model.User

	query := "SELECT password FROM users WHERE username = $1;"
	row := db.QueryRow(query, username)

	if err := row.Scan(&user.Username, &user.Password); err != nil {
		log.Println(err)
	}

	if user.Username == "" {
		return "user not found"
	}

	return user
}
