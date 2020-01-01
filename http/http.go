package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	DefaultAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"
)

func Get(apiUrl string, params url.Values) ([]byte, error) {
	// var Url *url.URL
	u, err := url.Parse(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("url parse error: \r\n %v", err)
	}
	// URLEncode
	u.RawQuery = params.Encode()

	client := &http.Client{}

	req, err := http.NewRequest("GET", u.String(), strings.NewReader(""))

	if err != nil {
		return nil, fmt.Errorf("http get error: %v", err)
	}

	req.Header.Set("User-Agent", DefaultAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func Post(apiUrl string, params string, headers map[string]string) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(params))
	if err != nil {
		return nil, fmt.Errorf("http post error: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", DefaultAgent)

	if headers != nil {
		for name, value := range headers {
			req.Header.Set(name, value)
		}
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("http post error:  %v", err)
	}

	return body, nil
}
