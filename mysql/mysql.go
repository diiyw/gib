package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Options struct {
	User     string
	Password string
	Host     string
	Port     string
	Db       string
}

var DefaultOptions = Options{
	User:     "root",
	Password: "",
	Host:     "127.0.0.1",
	Port:     "3306",
	Db:       "",
}

func (op Options) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", op.User, op.Password, op.Host, op.Port, op.Db)
}

type Mysql struct {
	db          *sql.DB
	builder     *Builder
	prevBuilder *Builder
}

func NewMysql(op Options) (*Mysql, error) {
	db, err := sql.Open("mysql", op.String())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Mysql{db: db, builder: new(Builder), prevBuilder: nil}, nil
}

func (my *Mysql) Table(table string) *Mysql {
	my.builder.table = Table(table, "")
	return my
}

func (my *Mysql) TableWithAlias(table, alias string) *Mysql {
	my.builder.table = Table(table, alias)
	return my
}

func (my *Mysql) LeftJoin(table, column, op, nextColumn string) *Mysql {
	my.builder.join = append(my.builder.join, LeftJoin(Table(table, ""), On(column, op, nextColumn)))
	return my
}

func (my *Mysql) Where(column, op string, bindValue interface{}) *Mysql {
	my.builder.cond = append(my.builder.cond, And(column, op, bindValue))
	return my
}

func (my *Mysql) WhereOr(column, op string, bindValue interface{}) *Mysql {
	my.builder.cond = append(my.builder.cond, Or(column, op, bindValue))
	return my
}

func (my *Mysql) WhereBetween(column, start, end string) *Mysql {
	my.builder.cond = append(my.builder.cond, AndBetween(column, []string{start, end}))
	return my
}

func (my *Mysql) WhereOrBetween(column, start, end string) *Mysql {
	my.builder.cond = append(my.builder.cond, OrBetween(column, []string{start, end}))
	return my
}

func (my *Mysql) GroupBy(by ...string) *Mysql {
	my.builder.groupBy = append(my.builder.groupBy, by...)
	return my
}

func (my *Mysql) OrderBy(by ...string) *Mysql {
	my.builder.orderBy = append(my.builder.orderBy, by...)
	return my
}

func (my *Mysql) Count(column string) (int, error) {

	my.builder.typ = Select("count(" + column + ")")

	sqlStr := my.Sql()
	stmt, err := my.db.Prepare(sqlStr)
	if err != nil {
		return 0, err
	}

	var count int
	if err := stmt.QueryRow(my.builder.bindValues).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (my *Mysql) Find(v []interface{}) error {

	my.builder.typ = Select("*")

	sqlStr := my.Sql()
	stmt, err := my.db.Prepare(sqlStr)
	if err != nil {
		return err
	}

	rows, err := stmt.Query(my.builder.bindValues)
	if err != nil {
		return err
	}
	var r = -1
	for rows.Next() {
		r++
		if err := rows.Scan(v[r]); err != nil {
			return err
		}
	}
	return nil
}

func (my *Mysql) FindOne(r interface{}) error {

	my.builder.typ = Select("*")

	sqlStr := my.Sql()
	stmt, err := my.db.Prepare(sqlStr)
	if err != nil {
		return err
	}

	row := stmt.QueryRow(my.builder.bindValues)
	return row.Scan(r)
}

func (my *Mysql) Insert(data map[string]interface{}) (int64, error) {

	my.builder.typ = Insert(data)

	result, err := my.exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (my *Mysql) ReplaceInto(data map[string]interface{}) (int64, error) {

	my.builder.typ = ReplaceInto(data)

	result, err := my.exec()
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (my *Mysql) Remove() (int64, error) {

	my.builder.typ = Delete()

	result, err := my.exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (my *Mysql) Update(data map[string]interface{}) (int64, error) {

	my.builder.typ = Update(data)

	result, err := my.exec()
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (my *Mysql) exec() (sql.Result, error) {
	sqlStr := my.Sql()
	stmt, err := my.db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(my.builder.bindValues)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (my *Mysql) Sql() string {
	return my.builder.make()
}

func (my *Mysql) Prepare(query string) (*sql.Stmt, error) {
	return my.db.Prepare(query)
}
