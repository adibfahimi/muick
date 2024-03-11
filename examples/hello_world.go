package main

import (
	"github.com/adibfahimi/muick"
)

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
		return c.SendString("You are logged in")
	})

	app.Listen(":3000")
}
