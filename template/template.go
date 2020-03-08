package template

import (
	"github.com/diiyw/gib/cache"
	"github.com/diiyw/gib/strings"
	"github.com/gobuffalo/packr/v2"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"path"
	"time"
)

// Template 模板
type Template struct {
	Box *packr.Box
	*template.Template
	// 缓存
	Cache  *cache.Cache
	Prefix string
	Files  []string
	Data   map[string]interface{}
}

// New 创建模板引擎
func New(box *packr.Box, options ...Option) *Template {
	t := &Template{
		Box:      box,
		Cache:    cache.New(),
		Prefix:   "template-",
		Files:    make([]string, 0),
		Template: template.New("template"),
	}

	for _, op := range options {
		op(t)
	}

	return t
}

// FuncMap 注册函数
func (tpl *Template) FuncMap(name string, fn interface{}) {
	tpl.Funcs(template.FuncMap{
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

	tpl.Template = tpl.Funcs(template.FuncMap{
		"pathContain": func(path, name string) string {
			if c.Request().RequestURI == path {
				return name
			}
			if path != "/" && strings.Has(c.Request().RequestURI, path) {
				return name
			}
			return ""
		},
		"paramsContain": func(param, value, name string) string {
			if strings.Format(param, strings.UrlDecode()) == value {
				return name
			}
			return ""
		},
		"datetime": func(t time.Time) string {
			return t.Format(strings.DateTimeFormat)
		},
		"html": func(c string) template.HTML {
			return template.HTML(c)
		},
	})

	if err := tpl.ParseComponent(tpl.Files...); err != nil {
		panic(err)
	}
	content, err := tpl.Box.FindString(name)
	if err != nil {
		return err
	}
	t, err := tpl.Clone()
	if err != nil {
		return err
	}
	if t, err = t.Parse(content); err != nil {
		return err
	}
	return t.Execute(w, data)
}

// ParseComponent 解析通用模板
func (tpl *Template) ParseComponent(files ...string) error {
	for _, file := range files {
		s, err := tpl.Box.FindString(file)
		if err != nil {
			return err
		}
		tpl.Template, err = tpl.Template.Parse(s)
		if err != nil {
			return err
		}
	}
	return nil
}

// LoadComponent 先加载后解析文件
func (tpl *Template) LoadComponent(files ...string) {
	tpl.Files = append(tpl.Files, files...)
}
