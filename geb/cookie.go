package geb

import "net/http"

func Cookie(opt ...CookieOption) *http.Cookie {

	cookie := new(http.Cookie)

	for _, op := range opt {
		op(cookie)
	}
	return cookie
}
