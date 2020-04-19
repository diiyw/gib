package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type Orm struct {
	*sqlx.DB
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
	db, err := sqlx.Open(orm.driver, orm.dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Orm{DB: db}, nil
}

type Model struct {
	ID       int       `db:"id"`
	CreateAt time.Time `db:"create_at"`
	UpdateAt time.Time `db:"update_at"`
	DeleteAt time.Time `db:"delete_at"`
	Status   int       `db:"status"`
}
