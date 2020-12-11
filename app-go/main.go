package main

import "github.com/gofiber/fiber/v2"

// Status is
type Status struct {
	Status   string
	Language string
}

// Hello is
type Hello struct {
	World string
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		data := Status{
			Status:   "running",
			Language: "Go",
		}

		return c.JSON(data)
	})

	app.Get("/hello", func(c *fiber.Ctx) error {

		data := Hello{
			World: c.Query("world"),
		}
		return c.JSON(data)
	})

	app.Listen(":4001")

}
