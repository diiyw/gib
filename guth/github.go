package guth

import (
	"encoding/json"
	"errors"
	"github.com/diiyw/gib/gttp"
	"regexp"
)

type GithubAuth struct {
	Config
}

// 获取登录地址
func (g *GithubAuth) GetRedirectURL(state string) string {
	return "https://github.com/login/oauth/authorize?client_id=" + g.AppID + "&redirect_uri=" + g.RedirectURL + "&state=" + state
}

// 获取token
func (g *GithubAuth) SetAccessToken(code string) error {
	resp, err := gttp.Do(
		gttp.Url("https://github.com/login/oauth/access_token?client_id=" + g.AppID + "&client_secret=" + g.AppKey + "&code=" + code + "&redirect_uri=" + g.RedirectURL),
	)
	if err != nil {
		return err
	}
	str := string(resp)
	if err, _ := regexp.MatchString("error", str); err {
		return errors.New("github error")
	}
	re, _ := regexp.Compile("access_token=(.*)&scope")
	access := re.FindStringSubmatch(str)
	if len(access) >= 2 {
		g.AccessToken = access[1]
	}
	return nil
}

//获取第三方用户信息
func (g *GithubAuth) GetUserInfo(v interface{}) error {
	resp, err := gttp.Do(
		gttp.Header(map[string]string{
			"Authorization": "token " + g.AccessToken,
		}),
		gttp.Url("https://api.github.com/user"),
	)
	if err != nil {
		return err
	}
	return json.Unmarshal(resp, v)
}
