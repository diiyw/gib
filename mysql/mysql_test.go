package mysql

import (
	"testing"
)

func TestNewMysql(t *testing.T) {
	DefaultDBConfig.Db = "test"
	my, err := NewMysql(DefaultDBConfig)
	if err != nil {
		t.Error(err)
		return
	}

	my = my.Table("test")

	my.Where("id", "=", 1)
	assert(my.Sql(), "SELECT * FROM test WHERE  id=?")

	my.Where("name", "=", "john")
	assert(my.Sql(), "SELECT * FROM test WHERE  id=? AND name=?")
}

func assert(sql, eq string) {
	if sql != eq {
		panic("except `" + eq + "` but got: `" + sql + "`")
	}
}
