package routes

import (
	_ "blogger/docs"
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"

	"blogger/internal/handlers"
	"blogger/internal/middlewares"
)

func Register(app *fiber.App) {
	// Auth
	app.Get("/register", handlers.RegisterForm)
	app.Post("/register", handlers.Register)
	app.Get("/login", handlers.LoginForm)
	app.Post("/login", handlers.Login)
	app.Get("/logout", handlers.Logout)

	// Web routes (JWT protected)
	app.Get("/", middlewares.JWTAuth(), handlers.ListPosts)
	app.Get("/post/:id", middlewares.JWTAuth(), handlers.ShowPost)
	app.Get("/new", middlewares.JWTAuth(), handlers.NewPostForm)
	app.Post("/create", middlewares.JWTAuth(), handlers.CreatePost)
	app.Get("/edit/:id", middlewares.JWTAuth(), handlers.EditPostForm)
	app.Post("/update/:id", middlewares.JWTAuth(), handlers.UpdatePost)
	app.Get("/delete/:id", middlewares.JWTAuth(), handlers.DeletePost)

	// API routes
	api := app.Group("/api", middlewares.JWTAuth())
	api.Get("/posts", handlers.GetPosts)
	api.Get("/posts/:id", handlers.GetPost)
	api.Post("/posts", handlers.CreatePostAPI)
	api.Put("/posts/:id", handlers.UpdatePostAPI)
	api.Delete("/posts/:id", handlers.DeletePostAPI)

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
}
