package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"

	api "dondu/API"
	ui "dondu/UI"
)

func main() {
	const PORT string = ":8080"
	server := fiber.New(fiber.Config{
		AppName: "Dondu",
	})

	// Rendering web app
	server.Use("*", filesystem.New(filesystem.Config{
		Root:       http.FS(ui.DistDir),
		PathPrefix: "dist",
		Browse:     true,
	}))

	// API routes
	api_group := server.Group("/api")
	api.Todo(api_group)

	server.Listen(PORT)
}
