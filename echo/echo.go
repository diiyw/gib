package echo

import (
	"github.com/labstack/echo/v4"
)

type App struct {
	*echo.Echo
	Addr string
}

var DefaultApp = &App{
	Addr: ":8080",
}

func Start(options ...Option) error {
	DefaultApp.Echo = echo.New()

	for _, option := range options {
		option(DefaultApp)
	}

	return DefaultApp.Start(DefaultApp.Addr)
}
