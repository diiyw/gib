package strings

import (
	"github.com/diiyw/gib/times"
	"github.com/labstack/echo/v4"
	"net/url"
	"strconv"
	"time"
)

func UrlParam(name string, c echo.Context) string {
	v := c.Param(name)
	str, _ := url.PathUnescape(v)
	return str
}

func ParseTime(format, name string, c echo.Context) time.Time {
	v := c.Param(name)
	return times.ParseTime(format, v)
}

func UnixTime(name string, c echo.Context) time.Time {
	v := c.Param(name)
	i, _ := strconv.ParseInt(v, 10, 64)
	return times.UnixTime(i)
}

func DateTime(name string, c echo.Context) time.Time {
	v := c.Param(name)
	return times.ParseTime(times.DateFormat, v)
}
