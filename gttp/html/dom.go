package html

import (
	"encoding/base64"
	"github.com/PuerkitoBio/goquery"
	"github.com/diiyw/h2md"
	"github.com/gocolly/colly/v2"
	"net/url"
)

type DOM struct {
	Name string `json:"name"`
	// Type text|html|img
	Type     string `json:"type"`
	Attr     string `json:"attr"`
	Selector string `json:"selector"`
}

func (dom DOM) getContent(e *goquery.Selection, c *colly.Collector, uri *url.URL) string {
	switch dom.Type {
	case "html":
		v, _ := e.Html()
		return v
	case "img":
		src, _ := e.Attr(dom.Attr)
		srcURL, _ := uri.Parse(src)
		return srcURL.String()
	case "base64text":
		return base64.StdEncoding.EncodeToString([]byte(e.Text()))
	case "base64html":
		v, _ := e.Html()
		return base64.StdEncoding.EncodeToString([]byte(v))
	case "base64link":
		src, _ := e.Attr(dom.Attr)
		srcURL, _ := uri.Parse(src)
		var content string
		c.OnResponse(func(response *colly.Response) {
			content = base64.StdEncoding.EncodeToString(response.Body)
		})
		_ = c.Visit(srcURL.String())
		return content
	case "markdown":
		v, _ := e.Html()
		md, _ := h2md.NewH2MD(v)
		return md.Text()
	default:
		return e.Text()
	}
}

func (dom DOM) getAttr(e *goquery.Selection) (v string) {
	v, _ = e.Attr(dom.Attr)
	return
}
