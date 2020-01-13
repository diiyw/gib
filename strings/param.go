package strings

import (
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

func FormParseTime(format, name string, c echo.Context) time.Time {
	v := c.Param(name)
	return ParseTime(format, v)
}

func FormUnixTime(name string, c echo.Context) time.Time {
	v := c.Param(name)
	i, _ := strconv.ParseInt(v, 10, 64)
	return UnixTime(i)
}

func DateTime(name string, c echo.Context) time.Time {
	v := c.Param(name)
	return ParseTime(strings.DateFormat, v)
}
