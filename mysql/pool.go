package mysql

import (
	"github.com/diiyw/gib/config"
	"github.com/diiyw/gib/errors"
	"github.com/jinzhu/gorm"
)

var connPool = make(map[string]*Mysql)

type Configs struct {
	mysql map[string]DBConfig
}

var (
	conf = Configs{
		mysql: make(map[string]DBConfig),
	}
)

func GetMysql(model string) *Mysql {
	if m, ok := connPool[model]; ok {
		return m
	}
	return nil
}

func Init() error {

	if err := config.YamlConfig("mysql", &conf.mysql); err != nil {
		return err
	}

	for name, c := range conf.mysql {
		var err error
		if _, ok := connPool[name]; !ok {
			if connPool[name], err = NewMysql(c); err != nil {
				return errors.Throw(
					errors.WrapString("Get mysql connection:", err),
				)
			}
			if c.Prefix != "" {
				gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
					return c.Prefix + defaultTableName
				}
			}
		}
	}

	return nil
}
