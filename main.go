package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"

	ui "dondu/UI"
)

func main() {
	const PORT string = ":3000"
	server := fiber.New(fiber.Config{
		AppName: "Dondu",
	})
	server.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(ui.DistDir),
		PathPrefix: "dist",
		Browse:     true,
	}))
	server.Listen(PORT)
}
