package echo

type Option func(app *App)

func WithAddr(addr string) Option {
	return func(app *App) {
		app.Addr = addr
	}
}

func WithActions(actions ...Action) Option {
	return func(app *App) {
		for _, action := range actions {
			app.GET(action.Group, action.Get, action.Middlewares...)
			app.POST(action.Group, action.Post, action.Middlewares...)
			app.DELETE(action.Group, action.Delete, action.Middlewares...)
			app.PUT(action.Group, action.Put, action.Middlewares...)
		}
	}
}
