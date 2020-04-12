package html

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"net/url"
	"strconv"
)

type Collector struct {
	Link      string           `json:"link"`
	UserAgent string           `json:"user_agent"`
	Groups    map[string][]DOM `json:"groups"`
}

func (collector *Collector) Query(options ...Option) (result []map[string]string, err error) {
	result = make([]map[string]string, 0)
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
				var ret = make(map[string]string, 0)
				for _, doc := range docs {
					func(dom DOM) {
						next := el.DOM.Find(dom.Selector)
						next.Each(func(i int, nextSel *goquery.Selection) {
							key := dom.Name
							if _, ok := ret[key]; ok {
								key += "_" + strconv.Itoa(i)
							}
							if dom.Attr != "" && dom.Type == "text" {
								ret[key] = dom.getAttr(nextSel.First())
								return
							}
							ret[key] = dom.getContent(nextSel.First(), c.Clone(), currentRequestURL)
						})
					}(doc)
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

func Query(options ...Option) (result []map[string]string, err error) {
	finder := new(Collector)
	return finder.Query(options...)
}
