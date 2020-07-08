package html

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"net/url"
)

type Collector struct {
	Link      string           `json:"link"`
	UserAgent string           `json:"user_agent"`
	Groups    map[string][]DOM `json:"groups"`
}

func (collector *Collector) Query(options ...Option) (result []map[string]interface{}, err error) {
	result = make([]map[string]interface{}, 0)
	collector.UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"

	for _, op := range options {
		op(collector)
	}

	c := colly.NewCollector(
		colly.UserAgent(collector.UserAgent),
		colly.AllowURLRevisit(),
	)

	var currentRequestURL *url.URL
	c.OnResponse(func(response *colly.Response) {
		currentRequestURL = response.Request.URL
	})

	for group, domes := range collector.Groups {
		func(tag string, docs []DOM) {
			c.OnHTML(tag, func(el *colly.HTMLElement) {
				var ret = make(map[string]interface{}, 0)
				for _, doc := range docs {
					nextRet := func(dom DOM) []string {
						next := el.DOM.Find(dom.Selector)
						var nextReturn = make([]string, 0)
						next.Each(func(i int, nextSel *goquery.Selection) {
							if dom.Attr != "" && dom.Type == "text" {
								nextReturn = append(nextReturn, dom.getAttr(nextSel.First()))
								return
							}
							nextReturn = append(nextReturn, dom.getContent(nextSel.First(), c.Clone(), currentRequestURL))
						})
						return nextReturn
					}(doc)
					if len(nextRet) != 1 {
						ret[doc.Name] = nextRet
					} else {
						ret[doc.Name] = nextRet[0]
					}
				}
				if len(ret) != 0 {
					result = append(result, ret)
				}
			})
		}(group, domes)
	}

	err = c.Visit(collector.Link)
	return
}

func Query(options ...Option) (result []map[string]interface{}, err error) {
	finder := new(Collector)
	return finder.Query(options...)
}
