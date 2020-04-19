package guth

type Option func(g *Guth)

func WithGithub(id, key, redirectURL string) Option {
	return func(g *Guth) {
		g.c.AppID = id
		g.c.AppKey = key
		g.c.RedirectURL = redirectURL
	}
}
