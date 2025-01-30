package main 

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gocurd/database"
)

func main() {
	database.Connect()

	fmt.Println("Successfully Connected")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		})
	})
	r.Run()
}