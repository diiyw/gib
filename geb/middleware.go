package geb

import (
	"errors"
	"fmt"
	"github.com/diiyw/gib/gog"
	"net/url"
	"runtime"
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

func SimpleKeyAuth(key string) MiddlewareFunc {
	return func(next HandlerFunc) HandlerFunc {
		return func(c Context) error {
			if c.FormValue("key") != key {
				return errors.New("key uncorrected")
			}
			return next(c)
		}
	}
}
