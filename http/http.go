package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Get(apiUrl string, params url.Values) ([]byte, error) {
	// var Url *url.URL
	u, err := url.Parse(apiUrl)
	if err != nil {
		return nil, fmt.Errorf("analytic url error: \r\n %v", err)
	}

	// URLEncode
	u.RawQuery = params.Encode()
	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("http get error: %v", err)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func Post(apiUrl string, params url.Values) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(""))
	if err != nil {
		return nil, fmt.Errorf("http post error: %v", err)
	}

	req.PostForm = params
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent:", "Ck 1.0")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("http post error:  %v", err)
	}

	return body, nil
}
