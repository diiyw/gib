package web

type InitOption func(app *App)

func Init(r InitOption) {
	r(app)
}
