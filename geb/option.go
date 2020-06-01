package geb

import (
	"github.com/diiyw/gib/gog"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
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

func CORS(config middleware.CORSConfig) Option {
	return func(app *App) {
		// CORS
		app.Use(middleware.CORSWithConfig(config))
	}
}

func Prometheus(name string, customMetricsList ...[]*prometheus.Metric) Option {
	return func(app *App) {
		p := prometheus.NewPrometheus(name, nil, customMetricsList...)
		p.Use(app.Echo)
	}
}

func LoadConfig(names ...string) Option {
	return func(app *App) {
		for _, name := range names {
			b, err := ioutil.ReadFile(ConfDir + name + ".yml")
			if err != nil {
				gog.Fatal(err)
			}
			app.Config[name] = b
		}
	}
}
