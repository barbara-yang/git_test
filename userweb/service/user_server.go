package service

import (
	"entry_task/common/module"
	"entry_task/userweb/client"
	"errors"
	"fmt"
)

// GetUser provide remote call getUser method server
func GetUser(user module.User) (module.User, error) {
	rpcConn, err := client.NewRPCConn()
	if err != nil {
		fmt.Print("rpcConn creat fail")
		return module.User{}, errors.New(fmt.Sprintf("remote caller creat err: %s", err))
	}
	var f func(module.User) (module.User, error)
	rpcConn.Call("getUser", &f)
	user, err = f(user)
	client.CloseRPCConn(rpcConn)
	return user, err
}
