package gttp

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/url"
	"time"
)

type Option func(c *Client) error

func Url(addr string) Option {
	return func(c *Client) error {
		u, err := url.Parse(addr)
		c.Request.Host = u.Host
		c.Request.URL = u
		return err
	}
}

func Method(m string) Option {
	return func(c *Client) error {
		c.Request.Method = m
		return nil
	}
}

func UserAgent(ua string) Option {
	return func(c *Client) error {
		c.Header.Set("User-Agent", ua)
		return nil
	}
}

func ContentType(t string) Option {
	return func(c *Client) error {
		c.Header.Set("Content-Type", t)
		return nil
	}
}

func Header(headers map[string]string) Option {
	return func(c *Client) error {
		for name, value := range headers {
			c.Header.Set(name, value)
			return nil
		}
		return nil
	}
}

func Values(values url.Values) Option {
	return func(c *Client) error {
		enc := values.Encode()
		if c.Request.Method == "GET" {
			c.Request.URL.RawQuery = enc
			return nil
		}
		body := []byte(enc)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		c.Request.ContentLength = int64(len(body))
		c.Request.GetBody = func() (io.ReadCloser, error) {
			r := bytes.NewReader(body)
			return ioutil.NopCloser(r), nil
		}
		return nil
	}
}

func Timeout(d time.Duration) Option {
	return func(c *Client) error {
		c.Timeout = d
		return nil
	}
}
