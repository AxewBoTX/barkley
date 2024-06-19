package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/axewbotx/barkley/web"
	"github.com/axewbotx/barkley/web/routes"
)

// index(/) route handler
func IndexHandler(c echo.Context) error {
	return web.RenderTemplTemplate(c, http.StatusOK, routes.Index_Page())
}

// /test route handler

func TestHandler(c echo.Context) error {
	return web.RenderTemplTemplate(c, http.StatusOK, routes.Test_Page())
}
