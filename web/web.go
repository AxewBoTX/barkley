package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

// set WriteHeader for response to the provided status_code, set the Content-Type to be sent as
// response to the request, and then render the templ component to HTML string
func RenderTemplTemplate(c echo.Context, status int, comp templ.Component) error {
	c.Response().Writer.WriteHeader(status)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	if template_render_err := comp.Render(c.Request().Context(), c.Response().Writer); template_render_err != nil {
		log.Error(
			"Failed To Render Response Template",
			"Error",
			template_render_err,
		)
		return c.String(http.StatusInternalServerError, "Failed To Render Response Template")
	}
	return nil
}
