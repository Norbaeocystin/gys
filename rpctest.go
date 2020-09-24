package main

import (
	"gys/pkg/gysrpc"
	"time"
)

func main() {
	go gysrpc.ServeRpc()
	time.Sleep(time.Second)
	gysrpc.CallClient()
}
