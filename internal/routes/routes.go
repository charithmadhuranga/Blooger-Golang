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

	// Web routes
	app.Get("/", middlewares.OptionalUser(), handlers.ListPosts)
	app.Get("/post/:id", middlewares.OptionalUser(), handlers.ShowPost)
	app.Get("/new", middlewares.JWTAuthRedirect(), handlers.NewPostForm)
	app.Post("/create", middlewares.JWTAuthRedirect(), handlers.CreatePost)
	app.Get("/edit/:id", middlewares.JWTAuthRedirect(), handlers.EditPostForm)
	app.Post("/update/:id", middlewares.JWTAuthRedirect(), handlers.UpdatePost)
	app.Get("/delete/:id", middlewares.JWTAuthRedirect(), handlers.DeletePost)

	// API routes
	// Public read
	apiPublic := app.Group("/api")
	apiPublic.Get("/posts", handlers.GetPosts)
	apiPublic.Get("/posts/:id", handlers.GetPost)
	// Authenticated write
	api := app.Group("/api", middlewares.JWTAuth())
	api.Post("/posts", handlers.CreatePostAPI)
	api.Put("/posts/:id", handlers.UpdatePostAPI)
	api.Delete("/posts/:id", handlers.DeletePostAPI)

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)
}
