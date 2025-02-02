// controller/controller.go
package controller

import (
	"github.com/gin-gonic/gin"
	"gocurd/services"
)

func UserController(r *gin.Engine) {
	r.POST("/create-post", services.CreatePost)
	r.GET("/get-post", services.GetPost)
	r.GET("/get-post/:id", services.GetPostByID)
	 // Add this missing route
	r.PUT("/update-post/:id", services.UpdatePost)
	r.DELETE("/delete-post/:id", services.DeletePost)
}