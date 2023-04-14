package repository

import (
	"database/sql"
	"log"

	"github.com/rafialariq/auth-jwt/model"
)

type UserRepo interface {
	LoginRepo(username string) any
}

type userRepo struct {
	db *sql.DB
}

func (u *userRepo) LoginRepo(username string) any {
	var user model.User

	query := "SELECT password FROM users WHERE username = $1;"
	row := u.db.QueryRow(query, username)

	if err := row.Scan(&user.Username, &user.Password); err != nil {
		log.Println(err)
	}

	if user.Username == "" {
		return "user not found"
	}

	return user
}
