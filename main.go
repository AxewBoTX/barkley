package main

import (
	"embed"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/axewbotx/barkley/handlers"
	"github.com/axewbotx/barkley/lib"
)

//go:embed all:public/lib
var LibDir embed.FS

func main() {
	// create the echo server object
	server := echo.New()

	// server static files from `public/lib` directory
	server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(LibDir),
		Root:       "/",
	}))

	// registering various routes using handler functions
	server.GET("/", handlers.IndexHandler)

	// this function runs when the whole main function i.e the server succesffully closes
	defer func() {
		server.Close()
		log.Info("Server closed successfully")
	}()

	// start the server and also handle the case of any error while shutting down
	if server_close_err := server.Start(lib.PORT); server_close_err != nil {
		log.Fatal("Server close error", "Error", server_close_err)
	}
}
