package rpc

import (
	"entry_task/common/config"
	"entry_task/rpcsrv"
	"fmt"
	"log"
)

var rpc *RPC

// RPC save the Singleton Object for server
type RPC struct {
	*rpcsrv.Server
}

// InitRPC create the connect for server
func InitRPC() {
	addr := fmt.Sprintf("%s:%d", config.RPCSERVER, config.RPCPORT)
	srv := rpcsrv.NewServer(addr)
	rpc = &RPC{srv}
}

// GetRPC will return an initRPC to server
func GetRPC() *RPC {
	if rpc == nil {
		log.Fatal("RPC init error")
	}
	return rpc
}
