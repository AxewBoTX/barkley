package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/axewbotx/barkley/core"
)

func SubTasks_Group_Handler(app *core.Application) {
	router := app.Server.Group("/api/subtasks")
	router.GET("/list", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "SubTask List")
	})
	router.POST("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Create SubTask")
	})
	router.PATCH("/:SubTaskID", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			fmt.Sprintf("Update SubTask With ID: %s", c.Param("SubTaskID")),
		)
	})
	router.DELETE("/:SubTaskID", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			fmt.Sprintf("Delete SubTask With ID: %s", c.Param("SubTaskID")),
		)
	})
}
