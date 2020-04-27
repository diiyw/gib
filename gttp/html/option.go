package html

type Option func(finder *Collector)

func Link(link string) Option {
	return func(c *Collector) {
		c.Link = link
	}
}

func UserAgent(ua string) Option {
	return func(c *Collector) {
		c.UserAgent = ua
	}
}

func WithDOM(dom DOM) Option {
	return func(c *Collector) {
		c.Groups["body"] = append(c.Groups["body"], dom)
	}
}
