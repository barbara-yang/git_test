package service

import (
	"entry_task/common/module"
	"entry_task/userweb/client"
	"log"
)

// Logout provide logout remote call server
func Logout(user module.User) error {
	rpcConn, err := client.NewRPCConn()
	if err != nil {
		log.Printf("rpcConn creat err: %s\n", err)
		return err
	}
	var f func(module.User) error
	rpcConn.Call("logout", &f)
	err = f(user)
	if err != nil {
		log.Printf("logout remote call err:%s", err)
		return err
	}
	client.CloseRPCConn(rpcConn)
	return nil
}
