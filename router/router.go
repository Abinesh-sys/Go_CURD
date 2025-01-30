package router

import (
	"github.com/gin-gonic/gin"
	"gocurd/controller"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(gin.Recovery())
	controller.UserController(r)
	return r
}