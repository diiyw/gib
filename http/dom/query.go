package dom

import (
	"github.com/gocolly/colly/v2"
	"net/url"
)

type Finder struct {
	Link      string           `json:"link"`
	UserAgent string           `json:"user_agent"`
	Groups    map[string][]DOM `json:"groups"`
}

func (finder *Finder) Query(options ...Option) (result []map[string]string, err error) {
	result = make([]map[string]string, 0)
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

	for group, domes := range finder.Groups {
		func(string, []DOM) {
			c.OnHTML(group, func(el *colly.HTMLElement) {
				var ret = make(map[string]string)
				for _, dom := range domes {
					sel := el.DOM.Find(dom.Selector)
					if dom.Attr != "" {
						ret[dom.Name] = dom.getAttr(sel.First())
						continue
					}
					ret[dom.Name] = dom.getContent(currentRequestURL, sel.First(), c.Clone())
				}
				result = append(result, ret)
			})
		}(group, domes)
	}

	err = c.Visit(finder.Link)
	return
}

func Query(options ...Option) (result []map[string]string, err error) {
	finder := new(Finder)
	return finder.Query(options...)
}
