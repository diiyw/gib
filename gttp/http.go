package gttp

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	*http.Client
	*http.Request
}

var defaultClient = new(Client)

func Do(options ...Option) ([]byte, error) {

	defaultClient.Client = new(http.Client)

	defaultClient.Request, _ = http.NewRequest("GET", "http://127.0.0.1/", nil)

	defaultClient.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")

	for _, option := range options {
		if err := option(defaultClient); err != nil {
			return nil, err
		}
	}

	resp, err := defaultClient.Do(defaultClient.Request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
