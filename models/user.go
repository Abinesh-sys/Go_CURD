package models

import (
	"gocurd/database"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID      uint            `json:"id" gorm:"primaryKey"`
	Title 	string 			`json:"title" validate:"required"`
	Body 	string 			`json:"body" validate:"required"`
	TagList pq.StringArray 	`json:"taglist" gorm:"type:text[]"`

}

func Migrate() {
	database.DB.AutoMigrate(&User{})
}