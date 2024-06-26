package main

import (
	"embed"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/axewbotx/barkley/core"
	"github.com/axewbotx/barkley/handlers"
	"github.com/axewbotx/barkley/handlers/api"
)

//go:embed all:public/lib
var LibDir embed.FS

func main() {
	// base application setup
	app := core.NewApplication(core.Config{
		Port: core.PORT,
	}).ConnectDatabase(gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	app.Database.HandleMigrations()

	// server static files from `public/lib` directory
	app.Server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(LibDir),
		Root:       "/",
	}))

	// route handlers
	app.Server.GET("/", handlers.IndexHandler)
	app.Server.GET("/test", handlers.TestHandler)
	// API route groups
	api.Projects_Group_Handler(app)
	api.Sections_Group_Handler(app)
	api.Tasks_Group_Handler(app)
	api.SubTasks_Group_Handler(app)

	// this function runs when the whole main function i.e the server successfully closes
	defer func() {
		app.Server.Close()
		log.Info("Server closed successfully")
	}()

	// start the server and also handle the case of any error while shutting down
	if server_close_err := app.Listen(); server_close_err != nil {
		log.Fatal("Server close error", "Error", server_close_err)
	}
}
