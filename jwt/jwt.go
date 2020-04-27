package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/diiyw/gib/gash"
	"github.com/diiyw/gib/geb"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var JWToken *jwt.Token

var key string

type MapClaims = jwt.MapClaims

func init() {
	if JWToken == nil {
		JWToken = jwt.New(jwt.SigningMethodHS256)
	}
	if key == "" {
		key = gash.MD5(gash.UUID())
	}
}

func GetToken(options ...Option) (string, error) {
	for _, op := range options {
		op(JWToken)
	}
	// Generate encoded token and send it as response.
	t, err := JWToken.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return t, nil
}

func Middleware() echo.MiddlewareFunc {
	return middleware.JWT([]byte(key))
}

func GetMapClaims(c geb.Context) (MapClaims, error) {
	u := c.Get("user")
	user, ok := u.(*jwt.Token)
	if !ok {
		return nil, errors.New("token occupied")
	}
	return user.Claims.(MapClaims), nil
}
