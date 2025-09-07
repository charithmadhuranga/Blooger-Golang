package config

import (
	"log"
	"time"

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

	// Seed sample posts if empty
	var count int64
	DB.Model(&models.Post{}).Count(&count)
	if count == 0 {
		now := time.Now()
		DB.Create(&models.Post{Title: "Welcome to Go Blog", Content: "This is your first post.", CreatedAt: now})
		DB.Create(&models.Post{Title: "Getting Started", Content: "Use the New Post button to add content.", CreatedAt: now})
	}
}
