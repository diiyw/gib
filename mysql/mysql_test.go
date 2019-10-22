package mysql

import (
	"testing"
)

func TestNewMysql(t *testing.T) {
	DefaultOptions.Db = "test"
	my, err := NewMysql(DefaultOptions)
	if err != nil {
		t.Error(err)
		return
	}

	my = my.Table("test")

	my.Where("id=?")
	assert(my.BindData(1).Sql(), "SELECT * FROM test WHERE  id=?")

	my.Reset()
	my.Where("id=?", "AND name=?")
	assert(my.BindData(1).Sql(), "SELECT * FROM test WHERE  id=? AND name=?")
}

func assert(sql, eq string) {
	if sql != eq {
		panic("except `" + eq + "` but got: `" + sql + "`")
	}
}
