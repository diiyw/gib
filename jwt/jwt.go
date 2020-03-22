package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/diiyw/gib/hash"
	"github.com/diiyw/gib/web"
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
		key = hash.MD5(hash.UUID())
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

func GetMapClaims(c web.Context) (MapClaims, error) {
	u := c.Get("user")
	user, ok := u.(*jwt.Token)
	if !ok {
		return nil, errors.New("token occupied")
	}
	return user.Claims.(MapClaims), nil
}
