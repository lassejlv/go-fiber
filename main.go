package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/jet/v2"
	"github.com/joho/godotenv"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	godotenv.Load()

	users := []User{
		{
			Name: "Lasse",
			Age:  18,
		},
		{
			Name: "Alberte",
			Age:  18,
		},
	}

	engine := jet.New("./views", ".jet")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Enable logger, with custom format. (this is awesome)
	app.Use(logger.New(logger.Config{
		Format: "${status} - ${method} ${path}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {

		return c.Render("index", fiber.Map{
			"Name":  "lasse",
			"users": users,
		})
	})

	ADDR := os.Getenv("ADDR")
	app.Listen(ADDR)

}
