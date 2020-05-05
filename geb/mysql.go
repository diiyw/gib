package geb

import (
	"github.com/diiyw/gib/orm"
)

var conn = make(map[string]*orm.Orm, 0)

func DB(name string) *orm.Orm {
	if o, ok := conn[name]; ok {
		return o
	}
	var dbConf map[string]string
	if err := Config("mysql", &dbConf); err != nil {
		panic(err)
	}
	db, err := orm.Open(
		orm.DSN(dbConf[name]),
	)
	if err != nil {
		panic(err)
	}
	conn[name] = db
	return conn[name]
}
