package database

import (
	"gorm.io/gorm"
)

// DB gorm.DB 对象
var DB *gorm.DB

const (
	Mysql = iota
	Sqlite
	Pgsql
)

func Connect(name ...string) *gorm.DB {
	if len(name) == 0 {
		return ConnectMysql()
	}
	return DB
}
