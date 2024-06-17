package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Projects_Group_Handler(router *echo.Group) {
	router.GET("/list", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Project List")
	})
}
