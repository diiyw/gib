package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func YamlConfig(name string, conf interface{}) error {

	dir := os.Getenv("APP_PATH")
	if dir == "" {
		dir = "conf/"
	} else {
		dir += "conf/"
	}

	_ = os.Setenv("CONFIG_PATH", dir)

	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = "dev"
	}

	b, err := ioutil.ReadFile(dir + name + "." + mode + ".yml")
	if err == nil {
		if err := yaml.Unmarshal(b, conf); err != nil {
			return err
		}
	}
	return err

}
