package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitostamatti/notes-app-golang/app/models"
	"github.com/vitostamatti/notes-app-golang/app/utils"
)

func GetUsers(c *gin.Context) {
	users := models.GetUsers()
	c.JSON(http.StatusOK, users)
}

type UserResponse struct {
	ID       uint
	Username string
}

func CreateUser(c *gin.Context) {
	var reqUser models.User

	if err := c.BindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resUser, err := reqUser.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resUser)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := models.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")

	reqUser, err := models.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	c.BindJSON(&reqUser)

	ressUser, err := reqUser.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user could not be updated"})
		return
	}
	c.JSON(http.StatusOK, ressUser)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user, err := models.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CurrentUser(c *gin.Context) (models.User, error) {

	id, err := utils.ExtractTokenID(c)

	if err != nil {
		return models.User{}, err
	}

	user, err := models.GetUserById(fmt.Sprint(id))

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func CurrentSuperUser(c *gin.Context) (models.User, error) {
	user, err := CurrentUser(c)
	if err != nil {
		return user, err
	}

	if !user.IsSuperuser {
		err = fmt.Errorf("user is not superuser")
		return user, err
	}

	return user, err
}

func GetCurrentUser(c *gin.Context) {
	user, err := CurrentUser(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}
	var resUser UserResponse
	resUser.ID = user.ID
	resUser.Username = user.Username
	c.JSON(http.StatusOK, resUser)
}
