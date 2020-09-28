package gonf

import (
	"github.com/diiyw/gib/gache"
	"github.com/ghodss/yaml"
	"io/ioutil"
)

var ymlConfigCache = gache.New()

func Yaml(name string, v interface{}) error {
	if ymlConfigCache.Exits(name) {
		return yaml.Unmarshal(ymlConfigCache.Get(name).([]byte), v)
	}
	b, err := ioutil.ReadFile(name + ".yml")
	if err != nil {
		return err
	}
	ymlConfigCache.Set(name, b)
	return yaml.Unmarshal(b, v)
}
