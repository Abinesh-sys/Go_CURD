package services

import (
	"gocurd/database"
	"gocurd/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func DeletePost(c *gin.Context) {
    // Retrieve the post ID from the URL parameter
    postID := c.Param("id")

    // Find the post by its ID
    var user models.User
    result := database.DB.First(&user, postID)

    // Check if the post exists
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        }
        return
    }

    // Delete the post from the database
    if err := database.DB.Delete(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
        return
    }

    // Respond with success
    c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}