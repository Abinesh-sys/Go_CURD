package router

import (
	"github.com/gin-gonic/gin"
	"gocurd/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Serve static files (CSS, JS, images)
	r.Static("/static", "./views/static")

	// Serve HTML file for posts
	r.GET("/", func(c *gin.Context) {
		c.File("./views/index.html")
	})

	controller.UserController(r)

	return r
}
