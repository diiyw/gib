package orm

import (
	"strconv"
	"strings"
)

type Option func(o *Orm)

func Driver(d string) Option {
	return func(o *Orm) {
		o.driver = d
	}
}

func Auth(username, password string) Option {
	return func(o *Orm) {
		o.dsn = strings.Replace(o.dsn, "root", username, -1)
		o.dsn = strings.Replace(o.dsn, "password", password, -1)
	}
}

func Addr(host string, port int) Option {
	return func(o *Orm) {
		o.dsn = strings.Replace(o.dsn, "localhost", host, -1)
		o.dsn = strings.Replace(o.dsn, "3306", strconv.Itoa(port), -1)
	}
}

func Database(db string) Option {
	return func(o *Orm) {
		o.dsn = strings.Replace(o.dsn, "test", db, -1)
	}
}

func Charset(s string) Option {
	return func(o *Orm) {
		o.dsn = strings.Replace(o.dsn, "utf8", s, -1)
	}
}

func Socket(sock string) Option {
	return func(o *Orm) {
		o.dsn = strings.Replace(o.dsn, "localhost:3306", sock, -1)
		o.dsn = strings.Replace(o.dsn, "tcp", "unix", -1)
	}
}

func DSN(dsn string) Option {
	return func(o *Orm) {
		o.dsn = dsn
	}
}