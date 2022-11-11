package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitostamatti/notes-app-golang/app/utils"
	"golang.org/x/crypto/bcrypt"

	"github.com/vitostamatti/notes-app-golang/app/models"
)

func Login(c *gin.Context) {

	var reqUser models.User

	if err := c.BindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := LoginCheck(reqUser.Username, reqUser.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u, err := models.GetUserByUsername(username)

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}
