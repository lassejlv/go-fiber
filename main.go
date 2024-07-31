package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/jet/v2"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	engine := jet.New("./views", ".jet")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Middleware
	app.Static("/", "./public")
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {

		return c.Render("index", fiber.Map{
			"Name":  "lasse",
			"Title": "lasse",
		})
	})

	ADDR := os.Getenv("ADDR")
	log.Fatal(app.Listen(ADDR))

}
