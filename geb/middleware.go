package geb

import (
	"fmt"
	"github.com/diiyw/gib/gog"
	"github.com/labstack/echo/v4"
	"net/url"
	"runtime"
	"strconv"
)

// Recover panic log
func Recover(next HandlerFunc) HandlerFunc {
	return func(c Context) error {

		defer func() {
			if r := recover(); r != nil {
				err, _ := r.(error)
				stack := make([]byte, 4<<10)
				length := runtime.Stack(stack, false)

				panicString := fmt.Sprintf("Error: %v \n %s", err, stack[:length])
				gog.Stdout(panicString)
				gog.File(panicString)

			}
		}()
		return next(c)
	}
}

// RequestLogger record request log
func RequestLogger(next HandlerFunc) HandlerFunc {
	return func(c Context) error {
		params := c.QueryParams()
		formParams, _ := c.FormParams()
		gog.File(map[string]interface{}{
			"ip":   c.RealIP(),
			"uri":  c.Request().RequestURI,
			"data": append([]url.Values{}, params, formParams),
		})
		return next(c)
	}
}

func KeyAuth(key string, f func(c Context) error) MiddlewareFunc {
	return func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			if c.FormValue("key") != key {
				return f(c)
			}
			return next(c)
		}
	}
}

func FormatPageParams(next HandlerFunc) HandlerFunc {
	return func(c echo.Context) error {
		page, _ := strconv.ParseInt(c.FormValue("page"), 10, 64)
		if page == 0 {
			page = 1
		}
		c.Set("page", int(page))
		limit, _ := strconv.ParseInt(c.FormValue("limit"), 10, 64)
		if limit == 0 {
			limit = 10
		}
		c.Set("limit", int(limit))
		offset := (page - 1) * limit
		c.Set("offset", int(offset))
		return next(c)
	}
}
