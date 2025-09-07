package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"blogger/internal/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	DB.AutoMigrate(&models.Post{}, &models.User{})
}
