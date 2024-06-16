//go:generate go run github.com/steebchen/prisma-client-go db push

package main

import (
	"context"
	"embed"
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/axewbotx/barkley/db"
	"github.com/axewbotx/barkley/handlers"
	"github.com/axewbotx/barkley/lib"
)

//go:embed all:public/lib
var LibDir embed.FS

func main() {
	server := echo.New()
	ctx := context.Background()
	DB := db.NewClient()
	if err := DB.Prisma.Connect(); err != nil {
		log.Error("Failed to connect database client", "Error", err)
	}
	createdPost, err := DB.Project.CreateOne(
		db.Project.Name.Set("Barkley"),
	).Exec(ctx)
	if err != nil {
		log.Error("Failed to create project", "Error", err)
	}

	result, _ := json.MarshalIndent(createdPost, "", "  ")
	log.Printf("created post: %s\n", result)

	// server static files from `public/lib` directory
	server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(LibDir),
		Root:       "/",
	}))

	// registering various routes using handler functions
	server.GET("/", handlers.IndexHandler)

	// this function runs when the whole main function i.e the server successfully closes
	defer func() {
		if err := DB.Prisma.Disconnect(); err != nil {
			panic(err)
		}
		server.Close()
		log.Info("Server closed successfully")
	}()

	// start the server and also handle the case of any error while shutting down
	if server_close_err := server.Start(lib.PORT); server_close_err != nil {
		log.Fatal("Server close error", "Error", server_close_err)
	}
}
