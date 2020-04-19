package guth

type Guth struct {
	c *Config
}

func NewGuth(options ...Option) *Guth {
	g := &Guth{}
	for _, opt := range options {
		opt(g)
	}
	return g
}
