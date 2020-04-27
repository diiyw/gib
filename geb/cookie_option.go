package geb

import "net/http"

type CookieOption func(cookie *http.Cookie)

func Set(name, value string) CookieOption {
	return func(cookie *http.Cookie) {
		cookie.Name = name
		cookie.Value = value
	}
}
