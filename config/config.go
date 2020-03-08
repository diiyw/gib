package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func YamlConfig(name string, conf interface{}) error {
	b, err := ioutil.ReadFile(ConfDir + name + ".yml")
	if err == nil {
		if err := yaml.Unmarshal(b, conf); err != nil {
			return err
		}
	}
	return err
}
