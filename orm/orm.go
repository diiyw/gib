package orm

import (
	"github.com/diiyw/gib/gache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Orm struct {
	*gorm.DB
	dsn string
}

func Open(options ...Option) (orm *Orm, err error) {
	orm = new(Orm)
	orm.dsn = "root:password@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local"
	for _, op := range options {
		op(orm)
	}
	db, err := gorm.Open(mysql.Open(orm.dsn), &gorm.Config{})
	orm.DB = db
	if err != nil {
		return nil, err
	}
	return orm, nil
}

var connCache *gache.Cache

func init() {
	if connCache == nil {
		connCache = gache.New()
	}
}

func Mysql(name string, options ...Option) *Orm {
	if connCache != nil && connCache.Exits(name) {
		return connCache.Get(name).(*Orm)
	}
	o, err := Open(options...)
	if err != nil {
		panic(err)
	}
	connCache.Set(name, o)
	return o
}

// EqualNotEmptyCondition 不为空就添加等于条件
func (orm *Orm) EqualNotEmptyCondition(cond map[string]string) *gorm.DB {
	var db = orm.DB
	for field, value := range cond {
		if value != "" {
			db = db.Where(field+"=?", value)
		}
	}
	return db
}

// Paginate 分页
func (orm *Orm) Paginate(page, perPage int) *gorm.DB {
	var db = orm.DB
	return db.Offset((page - 1) * perPage).Limit(perPage)
}