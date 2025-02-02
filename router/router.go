package router

import (
	"github.com/gin-gonic/gin"
	"gocurd/controller"
)

// SetupRouter initializes routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Static files and templates
	r.Static("/static", "./views/static")
	r.LoadHTMLGlob("views/*.html")

	// Base routes
	r.GET("/", func(c *gin.Context) {
		c.File("./views/index.html")  // Home page
	})

	// Route to display all posts
	r.GET("/allpost", func(c *gin.Context) {
		c.File("./views/allpost.html")  // Page showing all posts
	})

	// Route to view a single post by ID
	r.GET("/post/:id", func(c *gin.Context) {
    // Here, you should fetch the post by ID and render its details
    c.HTML(200, "post.html", gin.H{
        "postID": c.Param("id"),
    })
})


	// Route to edit a post by ID
	r.GET("/edit-post/:id", func(c *gin.Context) {
		c.File("./views/update.html")  // Page to edit a post
	})

	// Register controller routes for post actions
	controller.UserController(r)

	return r
}
