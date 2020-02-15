package echo

import "github.com/labstack/echo/v4"

type Action struct {
	Group       string
	Middlewares []echo.MiddlewareFunc
}

func (action *Action) Get(e echo.Context) error {
	return nil
}

func (action *Action) Post(e echo.Context) error {
	return nil
}

func (action *Action) Put(e echo.Context) error {
	return nil
}

func (action *Action) Delete(e echo.Context) error {
	return nil
}
