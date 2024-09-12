package db

import (
	"database/sql"
	"entry_task/common/config"
	"fmt"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *Database

// Database Object contain db
type Database struct {
	*sql.DB
}

// InitMysql initial Mysql database
func InitMysql() {
	d, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUSERNAME, config.DBPASSWORD, config.MYSQLSERVER, config.MYSQLPORT, config.DATABASE))
	if err != nil {
		log.Fatalf("db create err:%s\n", err)
	}
	d.SetMaxOpenConns(config.DBMAXOPENCONN)
	d.SetMaxIdleConns(config.DBMAXIDLECONN)
	d.SetConnMaxLifetime(config.DBMAXLIFETIME)
	//d.LogMode(false)
	if err := d.Ping(); err != nil {
		log.Fatal("mysql connect fail")
	}
	db = &Database{d}
}

// GetMysqlClient return the database client to dao
func GetMysqlClient() *Database {
	if db == nil {
		log.Fatal("Mysql init error.")
	}
	return db
}
