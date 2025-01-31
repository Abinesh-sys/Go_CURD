package controller

import (
	"github.com/gin-gonic/gin"
	"gocurd/services"
)


func UserController (r *gin.Engine) {

	
	r.POST("/create-post",services.CreatePost)
	r.GET("/get-post",services.GetPost)

	r.DELETE("/delete-post/:id", services.DeletePost)

	
	

}