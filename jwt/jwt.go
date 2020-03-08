package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/diiyw/gib/hash"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var jwToken *jwt.Token

var key string

func init() {
	if jwToken == nil {
		jwToken = jwt.New(jwt.SigningMethodHS256)
	}
	if key == "" {
		key = hash.MD5(hash.UUID())
	}
}

func GetToken(options ...Option) (string, error) {
	for _, op := range options {
		op(jwToken)
	}
	// Generate encoded token and send it as response.
	t, err := jwToken.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return t, nil
}

func Middleware() echo.MiddlewareFunc {
	return middleware.JWT([]byte(key))
}
