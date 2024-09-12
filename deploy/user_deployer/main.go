package main

import (
	"entry_task/common/db"
	initer "entry_task/deploy/user_deployer/db"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	db.InitMysql()
	initer.TableInit()

	c := make(chan bool)
	go initer.CreateUser(c)

	// 捕获系统信号
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-c:
		// 插入完成
		db.GetMysqlClient().Close()
	case sig := <-sigs:
		// 捕获到中断信号
		println("Received signal:", sig)
		// 优雅关闭数据库连接
		db.GetMysqlClient().Close()
	}
}
