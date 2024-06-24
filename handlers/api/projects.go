package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/axewbotx/barkley/core"
)

func Projects_Group_Handler(app *core.Application) {
	router := app.Server.Group("/api/projects")

	// ----- fetch project list
	router.GET("/list", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Project List")
	})

	// ----- create project
	router.POST("/", func(c echo.Context) error {
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
