package configs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ConfigsAndRandomRoutes() *fiber.App {
	crs := cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, DELETE, PUT, OPTIONS",
		AllowHeaders: "*",
	})

	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // Limite de 10Mb para envio de arquivos
	})

	app.Use(crs)

	return app
}
