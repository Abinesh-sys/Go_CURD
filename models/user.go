package models

import (
	"gocurd/database"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Title string `json:"title" validate:"required"`
	Body  string `json:"title" validate:"required"`
	TagList pq.StringArray `json:"taglist" gorm:"type:text[]"`

}

func Migrate() {
	database.DB.AutoMigrate(&User{})
}