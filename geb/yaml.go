package geb

import (
	"github.com/diiyw/gib/gerr"
	"github.com/ghodss/yaml"
)

func Config(name string, v interface{}) error {
	if c, ok := app.Config[name]; ok {
		return yaml.Unmarshal(c, v)
	}
	return gerr.New(gerr.String(name + " config: not found."))
}
