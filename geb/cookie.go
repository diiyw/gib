package geb

import (
	"net/http"
	"time"
)

func Cookie(opt ...CookieOption) *http.Cookie {

	cookie := new(http.Cookie)

	cookie.Path = "/"

	cookie.Expires = time.Now().Add(time.Hour * 72)

	for _, op := range opt {
		op(cookie)
	}
	return cookie
}
