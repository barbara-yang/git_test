package main

import (
	"entry_task/userweb/client"
	"entry_task/userweb/config"
	"entry_task/userweb/router"
	"fmt"
	"net/http"
)

func main() {
	router.InitHTTPRouter()
	client.InitRPCPool()
	fmt.Println("start http server at:", config.SERVERPORT)
	http.ListenAndServe(":"+config.SERVERPORT, nil)
}
