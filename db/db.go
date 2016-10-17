// Copyright 2016 Nikugame. All Rights Reserved

package db

import (
	"database/sql"
	"strings"

	"github.com/nikgame/public/config"
	"github.com/nikgame/public/log"
	//use mysql database
	_ "github.com/go-sql-driver/mysql"
)

//DB 数据库连接池客户端
var DB *sql.DB

func init() {
	conf, _ := config.NewConfig("ini", "conf/settings.conf")
	path := strings.Join([]string{conf.DefaultString("Mysql::user", "root"), ":", conf.DefaultString("Mysql::password", "password"), "@tcp(", conf.DefaultString("Mysql::host", "localhost"), ":", conf.DefaultString("Mysql::port", "3306"), ")/", conf.DefaultString("Mysql::dbname", "test"), "?charset=utf8&loc=Asia%2FShanghai"}, "")
	log.Log(path)
	DB, _ = sql.Open("mysql", path)
	DB.SetMaxOpenConns(conf.DefaultInt("Mysql::open", 100))
	DB.SetMaxIdleConns(conf.DefaultInt("Mysql::idle", 10))
	if err := DB.Ping(); err != nil {
		log.Log("DB Ping ERROR: %v", err)
		panic(err)
	}
	log.Log("database connect ...")
}
