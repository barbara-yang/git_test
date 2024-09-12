package main

import (
	"entry_task/common/db"
	"entry_task/common/redis"
	"entry_task/common/rpc"
	"entry_task/usersrv/config"
	"entry_task/usersrv/router"
	"flag"
	"github.com/pkg/profile"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	flag.BoolVar(&config.TESTFLAG, "t", true, "open test mode")
	db.InitMysql()
	redis.InitRedis()
	rpc.InitRPC()
	router.InitRPCRouter()
	go rpc.GetRPC().Run()
	defer db.GetMysqlClient().Close()
	defer redis.GetRedisCli().Close()
	defer profile.Start(profile.NoShutdownHook).Stop()
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	//阻塞直到有信号传入
	_ = <-c

}
