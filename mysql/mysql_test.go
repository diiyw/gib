package mysql

import (
	"testing"
)

func TestNewMysql(t *testing.T) {
	DefaultDbConfig.Db = "test"
	_, err := NewMysql(DefaultDbConfig)
	if err != nil {
		t.Error(err)
		return
	}
}
