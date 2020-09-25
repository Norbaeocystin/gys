package main

import (
	"gys/pkg/gysrpc"
	"log"
	"time"
)

func main() {
	rh := gysrpc.CallClient()
	time.Sleep(5 * time.Second)
	r := gysrpc.CallClient2(rh)
	log.Println(r)
	rh2 := gysrpc.CallClient3(r)
	time.Sleep(20 * time.Second)
	gysrpc.CallClient4(rh2)
	log.Println("Done")
}