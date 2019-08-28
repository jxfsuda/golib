package sql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

const (
	Sqlite3    = "sqlite3"
	Mysql      = "mysql"
	Postgresql = "postgres"
)

//数据库连接封装
type DbConnection struct {
	Db *sqlx.DB
}

var initedDB *DbConnection

//获取数据库连接
func GetConnection() *DbConnection {
	return initedDB
}

//新建一个数据库连接
func NewDatabaseConnection(dbType string, dbPath string) *DbConnection {
	db, err := sqlx.Open(dbType, dbPath)
	if err != nil {
		log.Fatalf("连接数据库失败 %v", err)
	}
	db.SetConnMaxLifetime(time.Minute * 460)
	err = db.Ping()
	db.Rebind(":")
	initedDB = &DbConnection{db}
	return initedDB
}
