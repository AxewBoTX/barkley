package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/axewbotx/barkley/core"
)

func Sections_Group_Handler(app *core.Application) {
	router := app.Server.Group("/api/sections")
	router.GET("/list", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Section List")
	})
	router.POST("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Create Section")
	})
	router.PATCH("/:SectionID", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			fmt.Sprintf("Update Section With ID: %s", c.Param("SectionID")),
		)
	})
	router.DELETE("/:SectionID", func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			fmt.Sprintf("Delete Section With ID: %s", c.Param("SectionID")),
		)
	})
}
