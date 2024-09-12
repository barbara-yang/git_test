package client

import (
	"encoding/gob"
	"entry_task/common/module"
	"entry_task/rpcsrv"
	"entry_task/userweb/config"
	"fmt"
	"github.com/flyaways/pool"
	"log"
	"time"
)

// RPCPool tcp connect pool  singleton object to creat conn
var RPCPool *pool.TCPPool

// InitRPCPool create tcp pool
func InitRPCPool() {
	options := &pool.Options{
		InitTargets:  []string{fmt.Sprintf("%s:%s", config.RPCSERVER, config.RPCPORT)},
		InitCap:      config.INITCONN,
		MaxCap:       config.MAXCONN,
		DialTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 60,
		ReadTimeout:  time.Second * 60,
		WriteTimeout: time.Second * 60,
	}
	RPCPool, _ = pool.NewTCPPool(options)

	gob.Register(module.User{})

}

// NewRPCConn find conn from pool return to caller
func NewRPCConn() (rpcsrv.Client, error) {
	conn, err := RPCPool.Get()
	if err != nil {
		log.Printf("dial error: %v\n", err)
		return rpcsrv.Client{}, err
	}
	return *rpcsrv.NewClient(conn), nil
}

// CloseRPCConn put conn to pool as close
func CloseRPCConn(client rpcsrv.Client) {
	_ = RPCPool.Put(client.Conn())
}
