package geb

type AppOption func(app *App)

func Register(r AppOption) {
	r(app)
}
