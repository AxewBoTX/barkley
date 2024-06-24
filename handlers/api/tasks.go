package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/axewbotx/barkley/core"
)

func Tasks_Group_Handler(app *core.Application) {
	router := app.Server.Group("/api/tasks")
	router.GET("/list", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Task List")
	})
	router.POST("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Create Task")
	})
	router.PATCH("/:TaskID", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			fmt.Sprintf("Update Task With ID: %s", c.Param("TaskID")),
		)
	})
	router.DELETE("/:TaskID", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			fmt.Sprintf("Delete Task With ID: %s", c.Param("TaskID")),
		)
	})
}
