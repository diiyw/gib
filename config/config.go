package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func YamlConfig(name string, conf interface{}) error {

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

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
