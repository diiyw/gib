package echo

import (
	"encoding/json"
	"github.com/diiyw/gib/strings"
	"github.com/labstack/echo/v4"
	"net/url"
	"strconv"
	"time"
)

type Query struct {
	context echo.Context
}

func (q *Query) GetQuery(name string, types ...Type) {
	types[0](name, q.context)
}

type Type func(name string, e echo.Context)

func String(v *string) Type {
	return func(name string, e echo.Context) {
		param := e.Param(name)
		*v, _ = url.PathUnescape(param)
	}
}

func Int(v *int) Type {
	return func(name string, e echo.Context) {
		param := e.Param(name)
		*v, _ = strconv.Atoi(param)
	}
}

func Time(format string, t *time.Time) Type {
	return func(name string, e echo.Context) {
		v := e.Param(name)
		*t = strings.ParseTime(format, v)
	}
}

func UnixTime(t *time.Time) Type {
	return func(name string, e echo.Context) {
		v := e.Param(name)
		i, _ := strconv.ParseInt(v, 10, 64)
		*t = strings.UnixTime(i)
	}

}

func DateTime(format string, t *time.Time) Type {
	return func(name string, e echo.Context) {
		v := e.Param(name)
		*t = strings.ParseTime(format, v)
	}
}

func Json(v interface{}) Type {
	return func(name string, e echo.Context) {
		v := e.Param(name)
		_ = json.Unmarshal([]byte(v), v)
	}
}
