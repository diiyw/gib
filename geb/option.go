package geb

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4/middleware"
)

type Option func(app *App)

func Addr(addr string) Option {
	return func(app *App) {
		app.Addr = addr
	}
}

func Renderer() Option {
	return func(app *App) {
		app.Renderer = app.Template
	}
}

func CORS(allowOrigins, allowMethods []string) Option {
	return func(app *App) {
		// CORS
		app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: allowOrigins,
			AllowMethods: allowMethods,
		}))
	}
}

func Prometheus(name string, customMetricsList ...[]*prometheus.Metric) Option {
	return func(app *App) {
		p := prometheus.NewPrometheus(name, nil, customMetricsList...)
		p.Use(app.Echo)
	}
}
