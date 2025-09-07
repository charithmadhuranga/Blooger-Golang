package handlers

import (
	"net/http"
	"strconv"

	"blogger/internal/config"
	"blogger/internal/models"
	"github.com/gofiber/fiber/v2"
)

// GetPosts godoc
// @Summary List posts
// @Tags posts
// @Produce json
// @Success 200 {array} models.Post
// @Router /api/posts [get]
func GetPosts(c *fiber.Ctx) error {
	var posts []models.Post
	config.DB.Order("created_at desc").Find(&posts)
	return c.JSON(posts)
}

// GetPost godoc
// @Summary Get a post
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} models.Post
// @Failure 404 {object} fiber.Map
// @Router /api/posts/{id} [get]
func GetPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Post not found"})
	}
	return c.JSON(post)
}

// CreatePostAPI godoc
// @Summary Create post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body models.Post true "New post"
// @Success 201 {object} models.Post
// @Router /api/posts [post]
// @Security BearerAuth
func CreatePostAPI(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}
	config.DB.Create(post)
	return c.Status(201).JSON(post)
}

// UpdatePostAPI godoc
// @Summary Update post
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param post body models.Post true "Updated post"
// @Success 200 {object} models.Post
// @Failure 404 {object} fiber.Map
// @Router /api/posts/{id} [put]
// @Security BearerAuth
func UpdatePostAPI(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
	}
	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}
	config.DB.Save(&post)
	return c.JSON(post)
}

// DeletePostAPI godoc
// @Summary Delete post
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 204 {string} string "No content"
// @Router /api/posts/{id} [delete]
// @Security BearerAuth
func DeletePostAPI(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Post{}, id).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Delete failed"})
	}
	return c.SendStatus(204)
}
