package service

import (
	"entry_task/common/module"
	"entry_task/userweb/client"
	"log"
)

// EditNickname provide remote call server to handle edit nickname
func EditNickname(user module.User) (module.User, error) {
	rpcConn, err := client.NewRPCConn()
	if err != nil {
		log.Printf("rpcConn creat err: %s\n", err)
		return module.User{}, err
	}
	var f func(module.User) (module.User, error)
	rpcConn.Call("editNickname", &f)
	updateUser, err := f(user)
	if err != nil {
		log.Printf("editNickname remote call err:%s", err)
		return module.User{}, err
	}
	client.CloseRPCConn(rpcConn)
	return updateUser, nil
}

// EditUserProfilePic provide server to remote call editProfilePic
func EditUserProfilePic(user module.User) (module.User, error) {
	rpcConn, err := client.NewRPCConn()
	if err != nil {
		log.Printf("rpcConn creat err: %s\n", err)
		return module.User{}, err
	}
	var f func(module.User) (module.User, error)
	rpcConn.Call("editProfilePic", &f)
	updateUser, err := f(user)
	if err != nil {
		log.Printf("editProfilePic remote call err:%s", err)
		return module.User{}, err
	}
	client.CloseRPCConn(rpcConn)
	return updateUser, nil
}
