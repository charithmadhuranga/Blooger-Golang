package handlers

import (
	"blogger/internal/config"
	"blogger/internal/models"
	"blogger/internal/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func RegisterForm(c *fiber.Ctx) error {
	return c.Render("register", fiber.Map{"Title": "Register"})
}

func Register(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user := models.User{Username: username, Password: string(hash)}

	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(400).SendString("Username already exists")
	}
	return c.Redirect("/login")
}

func LoginForm(c *fiber.Ctx) error {
	return c.Render("login", fiber.Map{"Title": "Login"})
}

func Login(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return c.Status(401).SendString("Invalid credentials")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return c.Status(401).SendString("Invalid credentials")
	}

	// Create JWT
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return c.Status(500).SendString("Error generating token")
	}

	// Store token in HTTP-only cookie for web
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		HTTPOnly: true,
		Secure:   false, // set true in production with HTTPS
	})

	return c.Redirect("/")
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		HTTPOnly: true,
		MaxAge:   -1,
	})
	return c.Redirect("/login")
}
