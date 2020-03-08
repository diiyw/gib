package dom

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"net/url"
)

type Finder struct {
	Link      string `json:"link"`
	UserAgent string `json:"user_agent"`
	DOM       []DOM  `json:"dom"`
}

func (finder *Finder) Query(options ...Option) (result map[string][]string, err error) {
	result = make(map[string][]string, 0)
	finder.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"

	for _, op := range options {
		op(finder)
	}

	c := colly.NewCollector(
		colly.UserAgent(finder.UserAgent),
		colly.AllowURLRevisit(),
	)

	var currentRequestURL *url.URL
	c.OnResponse(func(response *colly.Response) {
		currentRequestURL = response.Request.URL
	})

	c.OnHTML("body", func(el *colly.HTMLElement) {
		for _, dom := range finder.DOM {
			selection := el.DOM.Find(dom.Selector)
			selection.Each(func(i int, sel *goquery.Selection) {
				if dom.Attr != "" {
					result[dom.Name] = append(result[dom.Name], dom.getAttr(sel))
					return
				}
				result[dom.Name] = append(result[dom.Name], dom.getContent(currentRequestURL, sel, c.Clone()))
			})
		}
	})

	err = c.Visit(finder.Link)
	return
}

func Query(options ...Option) (result map[string][]string, err error) {
	finder := new(Finder)
	return finder.Query(options...)
}
