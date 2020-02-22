package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var pool map[string]*Orm

type Orm struct {
	*gorm.DB
	driver string
	dsn    string
}

func Open(options ...Option) (orm *Orm, err error) {
	orm = new(Orm)
	orm.driver = "mysql"
	orm.dsn = "root:password@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"
	for _, op := range options {
		op(orm)
	}
	db, err := gorm.Open(orm.driver, orm.dsn)
	if err != nil {
		return nil, err
	}
	return &Orm{DB: db}, nil
}
