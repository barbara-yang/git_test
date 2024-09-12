package service

import (
	"entry_task/common/module"
	"entry_task/userweb/client"
	"log"
)

// Index provide index session userID Auth remote call 用户身份认证功能
func Index(user module.User) (bool, error) {
	rpcConn, err := client.NewRPCConn()
	if err != nil {
		log.Printf("rpcConn creat err: %s\n", err)
		return false, err
	}
	var f func(module.User) (bool, error)
	rpcConn.Call("auth", &f)
	return f(user)
}
