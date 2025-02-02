package services

import (
	"gocurd/database"
	"gocurd/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func UpdatePost(c *gin.Context) {
    // Get post ID from URL
    postID := c.Param("id")
    id, err := strconv.ParseUint(postID, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
        return
    }

    // Find existing post
    var post models.User
    if err := database.DB.First(&post, uint(id)).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
        return
    }

    // Bind update data
    var updatedData struct {
        Title   string `json:"title"`
        Body    string `json:"body"`
        TagList string `json:"tag_list"`
    }

    if err := c.ShouldBindJSON(&updatedData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Update fields
    post.Title = updatedData.Title
    post.Body = updatedData.Body
    if updatedData.TagList != "" {
        post.TagList = pq.StringArray(SplitTags(updatedData.TagList))
    }

    // Save changes
    if err := database.DB.Save(&post).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Post updated successfully",
        "post":    post,
    })
}

// Helper function to split `TagList` string into an array
func SplitTags(tagString string) []string {
	tags := strings.Split(tagString, ",")
	for i := range tags {
		tags[i] = strings.TrimSpace(tags[i]) // Remove extra spaces
	}
	return tags
}



