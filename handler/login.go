package handler

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rafialariq/auth-jwt/model"
	"github.com/rafialariq/auth-jwt/repository"
	// "github.com/rafialariq/auth-jwt/utils"
)

// type Handler interface {
// 	LoginHandler(c *gin.Context)
// }

// type handler struct {
// 	userRepo repository.UserRepo
// }

func LoginHandler(c *gin.Context, db *sql.DB, jwtKey any) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := repository.CheckCredential(user.Username, db)
	recUser := result.(model.User)

	if recUser.Username == user.Username && recUser.Password == user.Password {

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Username
		claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}

}
