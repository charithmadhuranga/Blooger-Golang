package handlers

import (
	"blogger/internal/config"
	"blogger/internal/models"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

// GetPosts godoc
// @Summary Get all posts
// @Description Retrieve all blog posts
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
// @Summary Get single post
// @Description Retrieve a single post by ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} models.Post
// @Failure 404 {object} models.ErrorResponse
// @Router /api/posts/{id} [get]
func GetPost(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(models.ErrorResponse{Error: "Post not found"})
	}
	return c.JSON(post)
}

// CreatePostAPI godoc
// @Summary Create post
// @Description Create a new blog post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body models.Post true "Post data"
// @Success 201 {object} models.Post
// @Failure 400 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /api/posts [post]
func CreatePostAPI(c *fiber.Ctx) error {
	post := new(models.Post)
	if err := c.BodyParser(post); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid data"})
	}
	config.DB.Create(post)
	return c.Status(201).JSON(post)
}

// UpdatePostAPI godoc
// @Summary Update post
// @Description Update existing post
// @Tags posts
// @Accept json
// @Produce json
// @Param id path int true "Post ID"
// @Param post body models.Post true "Post data"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /api/posts/{id} [put]
func UpdatePostAPI(c *fiber.Ctx) error {
	id := c.Params("id")
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(models.ErrorResponse{Error: "Post not found"})
	}
	if err := c.BodyParser(&post); err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Invalid data"})
	}
	config.DB.Save(&post)
	return c.JSON(post)
}

// DeletePostAPI godoc
// @Summary Delete post
// @Description Delete a post by ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {object} models.ErrorResponse
// @Security BearerAuth
// @Router /api/posts/{id} [delete]
func DeletePostAPI(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := config.DB.Delete(&models.Post{}, id).Error; err != nil {
		return c.Status(400).JSON(models.ErrorResponse{Error: "Delete failed"})
	}
	return c.SendStatus(204)
}
