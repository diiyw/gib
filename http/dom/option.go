package dom

type Option func(finder *Finder)

func Link(link string) Option {
	return func(finder *Finder) {
		finder.Link = link
	}
}

func UserAgent(ua string) Option {
	return func(finder *Finder) {
		finder.UserAgent = ua
	}
}

func WithDOM(dom DOM) Option {
	return func(finder *Finder) {
		finder.DOM = append(finder.DOM, dom)
	}
}
