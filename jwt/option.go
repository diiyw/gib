package jwt

import "github.com/dgrijalva/jwt-go"

type Option func(token *jwt.Token)

func Method(method jwt.SigningMethod) Option {
	return func(token *jwt.Token) {
		token.Method = method
	}
}

func Key(k string) Option {
	return func(t *jwt.Token) {
		key = k
	}
}

func Claim(c map[string]interface{}) Option {
	return func(token *jwt.Token) {
		// Set claims
		token.Claims = jwt.MapClaims(c)
	}
}
