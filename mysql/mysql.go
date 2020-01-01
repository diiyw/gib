package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
	Charset  string `yaml:"charset"`
	Prefix   string `yaml:"prefix"`
}

var DefaultDbConfig = DBConfig{
	User:     "root",
	Password: "",
	Host:     "127.0.0.1",
	Port:     "3306",
	Db:       "",
	Charset:  "utf8",
}

func (op DBConfig) String() string {

	if op.Charset == "" {
		op.Charset = "utf8"
	}
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", op.User, op.Password, op.Host, op.Port, op.Db, op.Charset)
}

type Mysql struct {
	*gorm.DB
}

func NewMysql(c DBConfig) (*Mysql, error) {
	db, err := gorm.Open("mysql", c.String())
	if err != nil {
		return nil, err
	}
	return &Mysql{DB: db}, nil
}
