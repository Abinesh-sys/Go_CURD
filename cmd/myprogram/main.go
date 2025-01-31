package main 

import (
	"fmt"
	"gocurd/database"
	"gocurd/models"
	"gocurd/router"
)

func main() {
	database.Connect()
	models.Migrate()

	fmt.Println("Successfully Connected")

	r :=router.SetupRouter()
	r.Run(":8081")

}