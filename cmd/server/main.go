package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"

	"blogger/internal/config"
	"blogger/internal/routes"
)

// @title GoFiber Blogger API
// @version 1.0
// @description Simple blogging platform API with Fiber, GORM, and SQLite.
// @host localhost:3000
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Initialize database
	config.InitDB()

	// Load HTML templates with Fiber
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "base",
	})

	// Static files
	app.Static("/static", "./static")

	// Register routes
	routes.Register(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
