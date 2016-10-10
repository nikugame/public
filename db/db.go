// Copyright 2016 Nikugame. All Rights Reserved

package db

import (
	"database/sql"
	"fmt"
	"public/config"
	"public/tools"
	//use mysql database
	_ "github.com/go-sql-driver/mysql"
)

//DB 数据库连接池客户端
var DB *sql.DB

func init() {
	var conf dbConf
	conf.init()
	DB, _ := sql.Open("mysql", conf.connPath())
	DB.SetMaxOpenConns(conf.open)
	DB.SetMaxIdleConns(conf.idle)
	if err := DB.Ping(); err != nil {
		fmt.Printf("[%s] DB Ping ERROR: %s", tools.TimeStampString(), err)
	}
	fmt.Printf("[%s] database connect ...", tools.TimeStampString())
}

type dbConf struct {
	host   string
	port   string
	user   string
	passwd string
	dbname string
	open   int
	idle   int
}

//init func() 加载配置文件，获取mysql链接参数
func (c *dbConf) init() {
	conf, _ := config.NewConfig("ini", "conf/settings.conf")
	c.host = conf.DefaultString("Mysql::host", "localhost")
	c.port = conf.DefaultString("Mysql::port", "3306")
	c.user = conf.DefaultString("Mysql::user", "root")
	c.passwd = conf.DefaultString("Mysql::password", "")
	c.dbname = conf.DefaultString("Mysql::dbname", "test")
	c.open = conf.DefaultInt("Mysql::open", 100)
	c.idle = conf.DefaultInt("Mysql::idle", 10)
}

//connPath func() string 生成mysql链接字符串
func (c *dbConf) connPath() string {
	auth := tools.StringJoin(":", c.user, c.passwd)
	path := tools.StringJoin(":", c.host, c.port)
	common := tools.StringJoin("?", c.dbname, "charset=utf8&loc=Asia%2FShanghai")
	return tools.StringJoin("", auth, "@", path, "/", common)
}
