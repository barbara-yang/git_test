package service

import (
	"entry_task/common/module"
	"entry_task/userweb/client"
	"errors"
	"fmt"
)

// Login provide Login remote call method
func Login(username string, password string) (module.User, error) {
	rpcConn, err := client.NewRPCConn()
	if err != nil {
		fmt.Print("rpcConn creat fail")
		return module.User{}, errors.New(fmt.Sprintf("remote caller creat err: %s", err))
	}
	var f func(username string, password string) (module.User, error)
	rpcConn.Call("login", &f) //rpc client
	user, err := f(username, password)
	client.CloseRPCConn(rpcConn)
	return user, err
}
