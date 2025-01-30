package services

import (
	"gocurd/database"
	"gocurd/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	var user models.User
	c.BindJSON(&user)

	//saving in db
	result:= database.DB.Create(&user)
	if result.Error !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})

}