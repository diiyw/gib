package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/diiyw/gib/gash"
	"github.com/diiyw/gib/geb"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
)

var JWToken *jwt.Token

var key string

type MapClaims = jwt.MapClaims

type Config struct {
	SigningMethod string
	SigningKey    interface{}
	Claims        jwt.Claims
	// Header lookup
	TokenLookup  string
	ContextKey   string
	ErrorHandler func(c geb.Context, err error) error
	keyFunc      func(t *jwt.Token) (interface{}, error)
}

const (
	AlgorithmHS256 = "HS256"
)

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

func Verify(config Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		if config.SigningKey == nil {
			config.SigningKey = []byte(key)
		}
		if config.SigningMethod == "" {
			config.SigningMethod = AlgorithmHS256
		}
		config.ContextKey = "user"
		config.Claims = jwt.MapClaims{}
		config.keyFunc = func(t *jwt.Token) (interface{}, error) {
			// Check the signing method
			if t.Method.Alg() != config.SigningMethod {
				return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
			}
			return config.SigningKey, nil
		}
		return func(c echo.Context) error {
			var err error
			token := c.Request().Header.Get(config.TokenLookup)
			jToken := new(jwt.Token)
			// Issue #647, #656
			if _, ok := config.Claims.(jwt.MapClaims); ok {
				jToken, err = jwt.Parse(token, config.keyFunc)
			} else {
				t := reflect.ValueOf(config.Claims).Type().Elem()
				claims := reflect.New(t).Interface().(jwt.Claims)
				jToken, err = jwt.ParseWithClaims(token, claims, config.keyFunc)
			}
			if err == nil && jToken.Valid {
				// Store user information from token into context.
				c.Set(config.ContextKey, jToken)
				return next(c)
			}
			if config.ErrorHandler != nil {
				return config.ErrorHandler(c, err)
			}
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "invalid or expired jwt",
				Internal: err,
			}
		}
	}
}

func GetData(c geb.Context) (map[string]interface{}, error) {
	u := c.Get("user")
	user, ok := u.(*jwt.Token)
	if !ok {
		return nil, errors.New("token occupied")
	}
	return user.Claims.(MapClaims), nil
}
