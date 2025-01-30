package services

import (
	"gocurd/database"
	"gocurd/models"
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {

	var user models.User
	c.BindJSON(&user)

	//validation
	validate :=validator.New()
	err := validate.Struct(user)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	//saving in db
	result:= database.DB.Create(&user)
	if result.Error !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})

}

func GetPost(c *gin.Context) {

	var user []models.User
	result:=database.DB.Find(&user)
	if result.Error !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": user})
}