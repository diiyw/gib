package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/diiyw/gib/gash"
	"github.com/diiyw/gib/geb"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
)

type Token = jwt.Token

type MapClaims = jwt.MapClaims

var JWToken *Token

var key string

type Config struct {
	SigningMethod string
	SigningKey    interface{}
	Claims        jwt.Claims
	// Header lookup
	TokenLookup    string
	ContextKey     string
	ErrorHandler   func(c geb.Context, err error) error
	SuccessHandler func(token *Token) error
	keyFunc        func(t *Token) (interface{}, error)
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

func Verify(token string, config Config) error {
	if config.SigningKey == nil {
		config.SigningKey = []byte(key)
	}
	if config.SigningMethod == "" {
		config.SigningMethod = AlgorithmHS256
	}
	config.ContextKey = "user"
	config.Claims = jwt.MapClaims{}
	config.keyFunc = func(t *Token) (interface{}, error) {
		// Check the signing method
		if t.Method.Alg() != config.SigningMethod {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return config.SigningKey, nil
	}
	var err error
	jwtToken := new(Token)
	if _, ok := config.Claims.(jwt.MapClaims); ok {
		jwtToken, err = jwt.Parse(token, config.keyFunc)
	} else {
		t := reflect.ValueOf(config.Claims).Type().Elem()
		claims := reflect.New(t).Interface().(jwt.Claims)
		jwtToken, err = jwt.ParseWithClaims(token, claims, config.keyFunc)
	}
	if err == nil && jwtToken.Valid {
		if config.SuccessHandler != nil {
			return config.SuccessHandler(jwtToken)
		}
	}
	return err
}

func VerifyMiddleware(config Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.SuccessHandler == nil {
				config.SuccessHandler = func(token *Token) error {
					// Store user information from token into context.
					c.Set(config.ContextKey, token)
					return next(c)
				}
			}
			err := Verify(c.Request().Header.Get(config.TokenLookup), config)
			if config.ErrorHandler == nil {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "invalid or expired jwt",
					Internal: err,
				}
			}
			return config.ErrorHandler(c, err)
		}
	}
}

func GetMapData(c geb.Context) (v map[string]interface{}, err error) {
	u := c.Get("user")
	user, ok := u.(*Token)
	if !ok {
		return nil, errors.New("token occupied")
	}
	if err := GetClaimsData(user.Claims, &v); err != nil {
		return nil, err
	}
	return
}

func GetClaimsData(claims jwt.Claims, v interface{}) error {
	b, _ := json.Marshal(claims)
	return json.Unmarshal(b, v)
}
