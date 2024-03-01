package main

import (
	"errors"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	api "dondu/API"
	ui "dondu/UI"
	"dondu/lib"
)

func main() {
	DB := lib.PrepareDatabase()
	lib.HandleMigrations(DB)

	defer func() {
		DB.Close()
	}()

	server := echo.New()

	// Rendering web app
	server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(ui.DistDir),
		Root:       "dist",
		Browse:     true,
	}))

	// API routes
	api_group := server.Group("/api")
	api.Todo(api_group, DB)

	if err := server.Start(lib.PORT); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
