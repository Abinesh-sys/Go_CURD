package services

import (
	"gocurd/database"
	"gocurd/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
// GetPostByID fetches a single post by its ID
func GetPostByID(c *gin.Context) {
	id := c.Param("id")

	var post models.User
	if result := database.DB.First(&post, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}




