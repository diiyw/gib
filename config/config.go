package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func YamlConfig(name string, conf interface{}) error {

	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		path = dir + "/conf"
	}

	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = "dev"
	}

	b, err := ioutil.ReadFile(path + "/" + name + "." + mode + ".yml")
	if err == nil {
		if err := yaml.Unmarshal(b, conf); err != nil {
			return err
		}
	}
	return err

}
