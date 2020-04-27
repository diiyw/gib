package guth

type Guth interface {
	GetRedirectURL(state string) string
	SetAccessToken(code string) error
	GetUserInfo(v interface{}) error
}

func Github(id, key, redirectURL string) Guth {
	return &GithubAuth{
		Config{
			AppID:       id,
			AppKey:      key,
			RedirectURL: redirectURL,
		},
	}
}
