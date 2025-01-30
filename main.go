package main 

import (
	"fmt"
	"gocurd/database"
)

func main() {
	database.Connect()

	fmt.Println("Successfully Connected")
}