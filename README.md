# muick

```
Warning: This project is for demonstration purposes only
and should not be used in production environments.
```

Muick is a lightweight HTTP framework for Go, inspired by the Fiber framework. Built without any dependencies.

## Features

- **Easy to use**: Muick provides an intuitive API that makes building web applications a breeze.
- **Robust routing**: Define your application's routes using simple, expressive syntax.
- **JSON support**: Effortlessly work with JSON data using built-in helpers.

## Example

```go
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
```

## Getting Started

To get started with Muick, you can simply clone the repository and start building your web applications. Make sure you have Go installed on your system.

```bash
git clone https://github.com/adibfahimi/muick.git
cd muick
go run examples/hello_world.go
```

Visit [Muick GitHub repository](https://github.com/adibfahimi/muick) for more information and documentation.
