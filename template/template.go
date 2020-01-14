package template

import (
	"html/template"
	"io"
	"path"

	"github.com/diiyw/gib/cache"
	"github.com/diiyw/gib/strings"
	"github.com/gobuffalo/packr/v2"
	"github.com/labstack/echo/v4"
)

// Template 模板
type Template struct {
	Box    *packr.Box
	Driver *template.Template
	// 缓存
	Cache  *cache.Cache
	Prefix string
	Files  []string
}

// New 创建模板引擎
func New(box *packr.Box, files ...string) *Template {
	t := &Template{Box: box, Cache: cache.New(), Prefix: "template-", Files: files}

	t.Driver = template.New("template")

	return t
}

// FuncMap 注册函数
func (tpl *Template) FuncMap(name string, fn interface{}) {
	tpl.Driver.Funcs(template.FuncMap{
		name: fn,
	})
}

// Render 模板渲染接口
func (tpl *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	var err error

	if name == "" || name == "/" {
		name = "/index.html"
	}

	ext := path.Ext(name)

	if ext == "" {
		name += ".html"
	}

	tpl.Driver = tpl.Driver.Funcs(template.FuncMap{
		"pathContain": func(path, name string) string {
			if c.Path() == path {
				return name
			}
			if path != "/" && strings.Has(c.Path(), path) {
				return name
			}
			return ""
		},
		"paramsContain": func(param, value, name string) string {
			if strings.UrlParam(param, c) == value {
				return name
			}
			return ""
		},
	})

	if err := tpl.ParseFiles(tpl.Files); err != nil {
		panic(err)
	}
	content, err := tpl.Box.FindString(name)
	if err != nil {
		return err
	}
	t, err := tpl.Driver.Clone()
	if err != nil {
		return err
	}
	if t, err = t.Parse(content); err != nil {
		return err
	}
	return t.Execute(w, data)
}

// ParseFiles 解析通用模板
func (tpl *Template) ParseFiles(files []string) error {
	for _, file := range files {
		s, err := tpl.Box.FindString(file)
		if err != nil {
			return err
		}
		tpl.Driver, err = tpl.Driver.Parse(s)
		if err != nil {
			return err
		}
	}
	return nil
}
