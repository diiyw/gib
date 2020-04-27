package geb

import (
	"github.com/diiyw/gib/gerr"
	"github.com/diiyw/gib/restful"
	"github.com/diiyw/gib/template"
	"github.com/gobuffalo/packr/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	*echo.Echo
	Addr     string
	Template *template.Template
}

type Context = echo.Context

type HandlerFunc = echo.HandlerFunc

type MiddlewareFunc = echo.MiddlewareFunc

var app = &App{
	Echo:     echo.New(),
	Addr:     ":8080",
	Template: template.New(packr.New("app", "template")),
}

func Start(options ...Option) error {

	app.HTTPErrorHandler = func(err error, c echo.Context) {
		var (
			code        int
			errorString interface{}
		)
		switch e := err.(type) {
		case *echo.HTTPError:
			code = e.Code
			errorString = e.Message
		case *gerr.Error:
			code = e.Code()
			errorString = e.Error()
		default:
			code = 500
			errorString = e.Error()
		}
		_ = c.JSON(code, restful.Error(errorString))
	}

	app.Use(middleware.Gzip())
	app.Use(RequestLogger, Recover)

	for _, option := range options {
		option(app)
	}

	return app.Start(app.Addr)
}
