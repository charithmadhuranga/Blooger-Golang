package handlers

import (
	"blogger/internal/config"
	"blogger/internal/models"
	"github.com/gofiber/fiber/v2"
)

func ListPosts(c *fiber.Ctx) error {
	var posts []models.Post
	config.DB.Order("created_at desc").Find(&posts)
	return c.Render("index", fiber.Map{"Title": "My Blog", "Posts": posts})
}

func ShowPost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(404).SendString("Post not found")
	}
	return c.Render("show", fiber.Map{"Title": post.Title, "Post": post})
}

func NewPostForm(c *fiber.Ctx) error {
	return c.Render("form", fiber.Map{"Title": "New Post"})
}

func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(400).SendString("Invalid form data")
	}
	config.DB.Create(post)
	return c.Redirect("/")
}

func EditPostForm(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(404).SendString("Post not found")
	}
	return c.Render("edit", fiber.Map{"Title": "Edit Post", "Post": post})
}

func UpdatePost(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(404).SendString("Post not found")
	}
	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).SendString("Invalid data")
	}
	config.DB.Save(&post)
	return c.Redirect("/")
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.Post{}, id)
	return c.Redirect("/")
}
