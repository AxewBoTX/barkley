package api

import (
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"

	"github.com/axewbotx/barkley/core"
	"github.com/axewbotx/barkley/db"
)

func Projects_Group_Handler(app *core.Application) {
	router := app.Server.Group("/api/projects")
	client := app.Database.Client

	// ----- fetch project list
	router.GET("/list", func(c echo.Context) error {
		if projects, projects_fetch_err := client.Project.FindMany().Take(5).Exec(app.Database.Context); projects_fetch_err != nil {
			log.Error("Failed to fetch database records", "Error", projects_fetch_err)
		} else {
			log.Info(fmt.Sprintf("Project List: %+v", projects))
		}
		return c.JSON(http.StatusOK, "Project List")
	})

	// ----- create project
	router.POST("/", func(c echo.Context) error {
		if _, project_create_err := client.Project.CreateOne(db.Project.Name.Set("dates")).Exec(app.Database.Context); project_create_err != nil {
			log.Error("Failed to create database record", "Error", project_create_err)
		} else {
			log.Info("Successfully created project!")
		}
		return c.JSON(http.StatusOK, "Create Project")
	})

	// ----- update project
	router.PATCH("/:ProjectID", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			fmt.Sprintf("Update Project With ID: %s", c.Param("ProjectID")),
		)
	})

	// ----- delete project
	router.DELETE("/:ProjectID", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			fmt.Sprintf("Delete Project With ID: %s", c.Param("ProjectID")),
		)
	})
}
