package main

import (
	"github.com/adibfahimi/muick"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	app := muick.New()

	app.Get("/", func(c *muick.Ctx) error {
		c.Append("Content-Type", "text/html")
		return c.SendString("<h1>Welcome to my website</h1>")
	})

	app.Get("/json", func(c *muick.Ctx) error {
		return c.JSON(muick.Map{
			"message": "Hello, World!",
		})
	})

	app.Post("/login", func(c *muick.Ctx) error {
		var data loginRequest
		if err := c.BodyParser(&data); err != nil {
			return c.Status(400).JSON(muick.Map{
				"error": "Invalid request",
			})
		}

		if data.Username != "admin" || data.Password != "admin" {
			return c.Status(401).JSON(muick.Map{
				"error": "Invalid credentials",
			})
		} else {
			return c.JSON(muick.Map{
				"message": "Logged in",
			})
		}
	})

	app.Listen(":3000")
}
