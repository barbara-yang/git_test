package router

import (
	"encoding/gob"
	"entry_task/common/module"
	"entry_task/common/rpc"
	"entry_task/usersrv/handler"
	"log"
)

// InitRPCRouter init rpc server route register method
func InitRPCRouter() {
	gob.Register(module.User{})
	rpcSrv := rpc.GetRPC()
	rpcSrv.Register("login", handler.Login)
	rpcSrv.Register("register", handler.Register)
	rpcSrv.Register("editNickname", handler.EditNickname)
	rpcSrv.Register("editProfilePic", handler.EditUserProfilePic)
	rpcSrv.Register("logout", handler.Logout)
	rpcSrv.Register("getUser", handler.GetUser)
	rpcSrv.Register("auth", handler.Auth)
	log.Println("service is running")
}
