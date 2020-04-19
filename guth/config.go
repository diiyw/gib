package guth

type Config struct {
	AppID       string
	AppKey      string
	RedirectURL string
	AccessToken string
}

func (c *Config) GetRedirectURL(state string) string { return "" }
func (c *Config) SetAccessToken(code string) error   { return nil }
func (c *Config) GetUserInfo(v interface{}) error    { return nil }
