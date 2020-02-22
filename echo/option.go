package echo

type Option func(app *App)

func Addr(addr string) Option {
	return func(app *App) {
		app.Addr = addr
	}
}

func Actions(actions ...Action) Option {
	return func(app *App) {
		for _, action := range actions {
			app.GET(action.Group, action.Get, action.Middleware...)
			app.POST(action.Group, action.Post, action.Middleware...)
			app.DELETE(action.Group, action.Delete, action.Middleware...)
			app.PUT(action.Group, action.Put, action.Middleware...)
		}
	}
}
