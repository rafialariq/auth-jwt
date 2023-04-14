package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafialariq/auth-jwt/config"
	"github.com/rafialariq/auth-jwt/model"
)

func main() {
	db := config.ConnectDb()
	defer db.Close()

}

func login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res := LoginRepo(user.Username)

}

// func loginRepos(username string) any {
// 	var user model.User

// 	query := "SELECT password FROM users WHERE username = $1;"
// 	row := u.db.QueryRow(query, username)

// 	if err := row.Scan(&user.Username, &user.Password); err != nil {
// 		log.Println(err)
// 	}

// 	if user.Username == "" {
// 		return "user not found"
// 	}

// 	return user
// }
