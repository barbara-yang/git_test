package config

import "time"

// const is global setting
const (
	//mysql

	DBUSERNAME    = "root"
	DBPASSWORD    = "12345678"
	MYSQLSERVER   = "127.0.0.1"
	MYSQLPORT     = 3306
	DATABASE      = "user_info"
	DBMAXOPENCONN = 100
	DBMAXIDLECONN = 100
	DBMAXLIFETIME = time.Minute * 10

	//redis

	REDISSREVER  = "127.0.0.1"
	REDISPORT    = 6379
	REDISMAXOPEN = 2048
	REDISMAXIDLE = 2048
	REDISRETRY   = 3
	REDISTIMEOUT = 5 * time.Second

	//rpc

	RPCSERVER = "127.0.0.1"
	RPCPORT   = 2333

	//加密运算盐

	SALT = "ARandomString"

	//默认图像路径

	DEFAULTPIC = "web/default/deafult.jpg"
)
