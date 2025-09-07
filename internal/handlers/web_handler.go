package handlers

import (
	"blogger/internal/config"
	"blogger/internal/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

func ListPosts(c *fiber.Ctx) error {
	var posts []models.Post
	config.DB.Order("created_at desc").Find(&posts)
	return c.Render("index", fiber.Map{"Title": "Home", "Posts": posts, "User": c.Locals("user")})
}

func ShowPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(404).SendString("Post not found")
	}
	return c.Render("show", fiber.Map{"Post": post, "User": c.Locals("user")})
}

func NewPostForm(c *fiber.Ctx) error {
	return c.Render("form", fiber.Map{"Title": "New Post", "User": c.Locals("user")})
}

func CreatePost(c *fiber.Ctx) error {
	title := c.FormValue("title")
	content := c.FormValue("content")
	post := models.Post{Title: title, Content: content, CreatedAt: time.Now()}
	config.DB.Create(&post)
	return c.Redirect("/")
}

func EditPostForm(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var post models.Post
	config.DB.First(&post, id)
	return c.Render("edit", fiber.Map{"Post": post, "User": c.Locals("user")})
}

func UpdatePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var post models.Post
	config.DB.First(&post, id)
	post.Title = c.FormValue("title")
	post.Content = c.FormValue("content")
	config.DB.Save(&post)
	return c.Redirect("/")
}

func DeletePost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	config.DB.Delete(&models.Post{}, id)
	return c.Redirect("/")
}
