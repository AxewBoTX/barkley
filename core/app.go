package core

import (
	"github.com/labstack/echo/v4"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Server *echo.Echo
}

func NewApplication(config Config) *Application {
	return &Application{
		Config: config,
		Server: echo.New(),
	}
}

func (app *Application) Listen() error {
	return app.Server.Start(app.Config.Port)
}
